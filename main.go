package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/dave/jennifer/jen"
)

func capitalize(s string) string {
	if s == "" {
		return s
	}

	b := []byte(s)
	if b[0] >= 'a' && b[0] <= 'z' {
		b[0] &^= 'a' - 'A'
	}

	return string(b)
}

func downcase(s string) string {
	if s == "" {
		return s
	}

	b := []byte(s)
	if b[0] >= 'A' && b[0] <= 'Z' {
		b[0] |= 'a' - 'A'
	}

	return string(b)
}

func camelToSnake(s string) string {
	runes := []rune(s)
	var b strings.Builder

	isUpper := func(r rune) bool { return r >= 'A' && r <= 'Z' }
	isLower := func(r rune) bool { return r >= 'a' && r <= 'z' }

	for i, r := range runes {
		if isUpper(r) && i > 0 {
			prev := runes[i-1]
			nextIsLower := i+1 < len(runes) && isLower(runes[i+1])
			if !isUpper(prev) || nextIsLower {
				b.WriteByte('_')
			}
		}

		if isUpper(r) {
			b.WriteRune(r + ('a' - 'A'))
			continue
		}

		b.WriteRune(r)
	}

	return b.String()
}

func parseHexOrDec(s string) int {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(strings.ToLower(s), "0x") {
		v, err := strconv.ParseInt(s[2:], 16, 64)
		if err != nil {
			log.Fatalf("invalid hex literal %q: %v", s, err)
		}
		return int(v)
	}

	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("invalid decimal literal %q: %v", s, err)
	}

	return int(v)
}

func renderHex(s string) jen.Code {
	return jen.Op(fmt.Sprintf("%#0X", parseHexOrDec(s)))
}

func parseCondition(cond string) *jen.Statement {
	cond = strings.TrimSpace(cond)

	if parts := strings.SplitN(cond, " || ", 2); len(parts) == 2 {
		return parseCondition(parts[0]).Op("||").Add(parseCondition(parts[1]))
	}

	if parts := strings.SplitN(cond, " && ", 2); len(parts) == 2 {
		return parseCondition(parts[0]).Op("&&").Add(parseCondition(parts[1]))
	}

	return parseAtomCond(cond)
}

func parseAtomCond(cond string) *jen.Statement {
	if !strings.Contains(cond, "=") {
		field := strings.TrimSpace(cond)
		return jen.Id(field)
	}

	for _, op := range []string{"!=", "=="} {
		before, after, ok := strings.Cut(cond, op)
		if !ok {
			continue
		}

		field := strings.TrimSpace(before)
		value := parseHexOrDec(strings.TrimSpace(after))

		return jen.Id(field).Op(op).Lit(value)
	}

	log.Fatalf("cannot parse condition atom: %q", cond)
	return nil
}

func conditionFieldNames(cond string) []string {
	if cond == "" {
		return nil
	}

	atoms := strings.FieldsFunc(cond, func(r rune) bool {
		return r == '|' || r == '&'
	})

	var names []string
	for _, atom := range atoms {
		atom = strings.TrimSpace(atom)

		if !strings.Contains(atom, "=") {
			names = append(names, strings.TrimSpace(atom))
			continue
		}

		for _, op := range []string{"!=", "=="} {
			if before, _, ok := strings.Cut(atom, op); ok {
				names = append(names, strings.TrimSpace(before))
				break
			}
		}
	}

	return names
}

// Type represents the Go type of an encoding field
type Type uint

const (
	TypeBool Type = iota
	TypeUint8
	TypeUint16
	TypeUint32
	TypeRegister
	TypeExtension
	TypeShift
	TypeBitmask
)

// jenCode returns the jennifer AST node for the type
func (t Type) jenCode() *jen.Statement {
	switch t {
	case TypeBool:
		return jen.Bool()
	case TypeUint8:
		return jen.Uint8()
	case TypeUint16:
		return jen.Uint16()
	case TypeUint32:
		return jen.Uint32()
	case TypeRegister:
		return jen.Id("Register")
	case TypeExtension:
		return jen.Id("Extension")
	case TypeShift:
		return jen.Id("Shift")
	case TypeBitmask:
		return jen.Uint64()
	default:
		log.Fatalf("unknown type %d", t)
		return nil
	}
}

// setFunc returns the appropriate setter
func (t Type) setFunc() *jen.Statement {
	switch t {
	case TypeBool:
		return jen.Id("setBool")

	case TypeRegister:
		return jen.Id("setReg")

	case TypeBitmask:
		return jen.Id("setBitmask")

	default:
		return jen.Id("set").Types(t.jenCode())
	}
}

// getFunc returns the appropriate getter
func (t Type) getFunc() *jen.Statement {
	switch t {
	case TypeBool:
		return jen.Id("getBool")

	case TypeRegister:
		return jen.Id("getReg")

	case TypeBitmask:
		return jen.Id("getBitmask")

	default:
		return jen.Id("get").Types(t.jenCode())
	}
}

func (t *Type) UnmarshalJSON(data []byte) error {
	switch key := string(bytes.Trim(data, `"`)); key {
	case "bool":
		*t = TypeBool
	case "uint8":
		*t = TypeUint8
	case "uint16":
		*t = TypeUint16
	case "uint32":
		*t = TypeUint32
	case "reg":
		*t = TypeRegister
	case "ext":
		*t = TypeExtension
	case "shift":
		*t = TypeShift
	case "bitmask":
		*t = TypeBitmask
	default:
		return fmt.Errorf("unknown field type: %s", key)
	}
	return nil
}

// SchemaField is the resolved, fully typed representation of a single encoding
// field
type SchemaField struct {
	Name    string
	Type    Type
	Autoset bool
	Size    int
	Pos     int
}

// setCall emits setter for a single field
func (f SchemaField) setCall(recv, value jen.Code) *jen.Statement {
	return f.Type.setFunc().Call(recv, value, jen.Lit(f.Size), jen.Lit(f.Pos))
}

// getCall emits getter for a single field
func (f SchemaField) getCall(recv jen.Code) *jen.Statement {
	return f.Type.getFunc().Call(recv, jen.Lit(f.Size), jen.Lit(f.Pos))
}

// Schema holds the computed base instruction value and named field metadata
type Schema struct {
	Base   uint32
	Fields map[string]SchemaField
}

// mustField looks up a field by name or fatally exits if it is not found
func (s Schema) mustField(name string) SchemaField {
	f, ok := s.Fields[name]
	if !ok {
		log.Fatalf("encoding does not contain field %q", name)
	}

	return f
}

// fieldsFor resolves an ordered list of field names into schemaField values
func (s Schema) fieldsFor(names []string) []SchemaField {
	fields := make([]SchemaField, len(names))
	for i, name := range names {
		fields[i] = s.mustField(name)
	}

	return fields
}

// Field is the raw JSON representation of one bit-field in an encoding
type Field struct {
	Type    Type   `json:"type"`
	Name    string `json:"name"`
	Default string `json:"default"`
	Autoset *bool  `json:"autoset"`
	Size    int    `json:"size"`
	Pos     int    `json:"pos"`
}

func setBits(inst, op uint32, size, pos int) uint32 {
	mask := (uint32(1) << size) - 1
	if op > mask {
		log.Fatalf("expected %d-bit value at pos %d, got %#X", size, pos, op)
	}

	return (inst &^ (mask << pos)) | ((op & mask) << pos)
}

// Encoding describes the bit layout and API surface of an instruction category
type Encoding struct {
	Fields  []Field  `json:"fields"`
	Setters []string `json:"setters"`
	Getters []string `json:"getters"`
	Formats []Format `json:"formats"`
}

// buildSchema computes the constant base value and collects named field
// metadata
func (e Encoding) buildSchema() Schema {
	var base uint32
	fields := make(map[string]SchemaField)

	for _, f := range e.Fields {
		var def uint32
		if f.Default != "" {
			def = uint32(parseHexOrDec(f.Default))
		}

		if f.Name == "" {
			// Anonymous field: bake the constant into the base word
			base = setBits(base, def, f.Size, f.Pos)
			continue
		}

		autoset := true
		if f.Autoset != nil {
			autoset = *f.Autoset
		}

		fields[f.Name] = SchemaField{
			Name:    f.Name,
			Type:    f.Type,
			Autoset: autoset,
			Size:    f.Size,
			Pos:     f.Pos,
		}
	}

	return Schema{Base: base, Fields: fields}
}

// Format describes one disassembly output variant, optionally gated by a
// condition
type Format struct {
	Format    string   `json:"format"`
	Condition string   `json:"condition"`
	Arguments []string `json:"arguments"`
}

// neededLocals returns the ordered, deduplicated set of field names (excluding
// "mnemonic") that must be pre-fetched before evaluating formats and their
// conditions
func neededLocals(formats []Format) []string {
	var result []string

	seen := make(map[string]bool)
	add := func(name string) {
		if name == "mnemonic" || seen[name] {
			return
		}
		seen[name] = true
		result = append(result, name)
	}

	for _, f := range formats {
		for _, arg := range f.Arguments {
			add(arg)
		}
		for _, name := range conditionFieldNames(f.Condition) {
			add(name)
		}
	}

	return result
}

// Discriminator is a field/value pair used to identify a mnemonic from the
// encoding
type Discriminator struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

// Instruction represents one instruction variant within a category
type Instruction struct {
	Mnemonic       string          `json:"mnemonic"`
	Function       string          `json:"function"`
	Description    string          `json:"description"`
	Discriminators []Discriminator `json:"discriminators"`
}

func discriminatorLocals(instructions []Instruction, s Schema) []SchemaField {
	seen := make(map[string]bool)
	var fields []SchemaField
	for _, inst := range instructions {
		for _, d := range inst.Discriminators {
			if !seen[d.Field] {
				seen[d.Field] = true
				fields = append(fields, s.mustField(d.Field))
			}
		}
	}
	return fields
}

// matchCond builds the boolean expression that identifies this mnemonic
func (i Instruction) matchCond() *jen.Statement {
	var expr *jen.Statement
	for _, d := range i.Discriminators {
		cond := jen.Id(d.Field).Op("==").Add(renderHex(d.Value))
		if expr == nil {
			expr = cond
			continue
		}

		expr = expr.Op("&&").Add(cond)
	}

	return expr
}

// Category maps to one generated Go file and type
type Category struct {
	Name         string        `json:"name"`
	Base         string        `json:"base"`
	Encoding     Encoding      `json:"encoding"`
	Instructions []Instruction `json:"instructions"`
}

func (c Category) recv() jen.Code { return jen.Id("i") }
func (c Category) typ() jen.Code  { return jen.Id(c.Name) }

// recvParam generates the method receiver declaration: i ArithCarry
func (c Category) recvParam() jen.Code { return jen.Id("i").Id(c.Name) }

func (c Category) generate(pkg string) (*jen.File, error) {
	s := c.Encoding.buildSchema()
	file := jen.NewFile(pkg)

	c.emitTypeDecl(file)
	c.emitEncoders(file, s)
	c.emitGetters(file, s)
	c.emitMnemonicMethod(file, s)
	c.emitStringMethod(file, s)

	return file, nil
}

// emitTypeDecl emits: type ArithCarry Instruction
func (c Category) emitTypeDecl(file *jen.File) {
	file.Type().Id(c.Name).Id(c.Base)
	file.Line()
}

// emitEncoders generates one constructor function per mnemonic:
//
//	func ADC(rd, rn, rm Register) ArithCarry { ... }
func (c Category) emitEncoders(file *jen.File, s Schema) {
	setterFields := s.fieldsFor(c.Encoding.Setters)
	for _, m := range c.Instructions {
		c.emitEncoder(file, s, m, setterFields)
	}
}

func sfSetterArgs(setterFields []SchemaField) []jen.Code {
	var args []jen.Code
	for _, f := range setterFields {
		if f.Type == TypeRegister {
			args = append(args, jen.Id(f.Name))
		}
	}

	return args
}

func (c Category) emitEncoder(
	file *jen.File,
	schema Schema,
	inst Instruction,
	setterFields []SchemaField,
) {
	recv := c.recv()

	file.Commentf(
		"%s %s",
		inst.Function,
		downcase(inst.Description),
	)

	file.Func().
		Id(inst.Function).
		Params(collapseParams(setterFields)...).
		Add(c.typ()).
		BlockFunc(func(g *jen.Group) {
			g.Var().Add(recv, c.typ()).Op("=").Op(fmt.Sprintf("0x%08X", schema.Base))
			g.Line()

			if sf, ok := schema.Fields["sf"]; ok && sf.Autoset {
				sfArgs := sfSetterArgs(setterFields)
				if len(sfArgs) > 0 {
					callArgs := append([]jen.Code{recv}, sfArgs...)
					g.Add(recv).Op("=").Id("setSF").Call(callArgs...)
					g.Line()
				}
			}

			for _, f := range setterFields {
				g.Add(recv).Op("=").Add(f.setCall(recv, jen.Id(f.Name)))
			}
			g.Line()

			for _, d := range inst.Discriminators {
				f := schema.mustField(d.Field)
				g.Add(recv).Op("=").Add(
					f.setCall(recv, renderHex(d.Value)),
				)
			}
			g.Line()

			g.Return(recv)
		})

	file.Line()
}

// emitGetters generates one getter method per name in Encoding.Getters:
//
//	func (i ArithCarry) Rd() Register { return getReg(i, 5, 0) }
func (c Category) emitGetters(file *jen.File, s Schema) {
	for _, name := range c.Encoding.Getters {
		f := s.mustField(name)
		file.Func().
			Params(c.recvParam()).
			Id(capitalize(f.Name)).
			Params().
			Add(f.Type.jenCode()).
			Block(
				jen.Return(f.getCall(c.recv())),
			)
		file.Line()
	}
}

// emitMnemonicMethod generates the Mnemonic() string method:
//
//	func (i ArithCarry) Mnemonic() string {
//	    switch {
//	    case get[uint8](i, 1, 30) == 0 && get[uint8](i, 1, 29) == 0:
//	        return "ADC"
//	    ...
//	    }
//	    return "UNALLOCATED"
//	}
func (c Category) emitMnemonicMethod(file *jen.File, s Schema) {
	locals := discriminatorLocals(c.Instructions, s)

	file.Func().
		Params(c.recvParam()).
		Id("Mnemonic").Params().String().
		BlockFunc(func(g *jen.Group) {
			for _, f := range locals {
				g.Id(f.Name).Op(":=").Add(f.getCall(c.recv()))
			}
			g.Line()

			if len(locals) == 1 {
				c.emitValueSwitch(g, locals[0])
				g.Line()
				g.Return(jen.Lit("UNALLOCATED"))
				return
			}

			c.emitBoolSwitch(g)
			g.Line()
			g.Return(jen.Lit("UNALLOCATED"))
		})

	file.Line()
}

func (c Category) emitValueSwitch(g *jen.Group, local SchemaField) {
	g.Switch(jen.Id(local.Name)).BlockFunc(func(g *jen.Group) {
		for _, inst := range c.Instructions {
			g.Case(renderHex(inst.Discriminators[0].Value))
			g.Return(jen.Lit(inst.Mnemonic))
		}
	})
}

func (c Category) emitBoolSwitch(g *jen.Group) {
	g.Switch().BlockFunc(func(g *jen.Group) {
		for _, inst := range c.Instructions {
			g.Case(inst.matchCond())
			g.Return(jen.Lit(inst.Mnemonic))
		}
	})
}

// emitStringMethod generates the String() string method.
//
// All field values referenced in arguments or conditions are pre-fetched into
// local variables to avoid redundant bit-field extractions and to keep the
// generated conditional expressions readable:
//
//	func (i ArithImm) String() string {
//	    mnemonic := i.Mnemonic()
//	    shift := get[uint8](i, 1, 22)
//	    imm   := get[uint16](i, 12, 10)
//	    rn    := getReg(i, 5, 5)
//	    rd    := getReg(i, 5, 0)
//	    if shift == 0x1 {
//				return fmt.Sprintf("%s %s, %s, #%d, LSL #12", mnemonic, rd, rn, imm)
//			}
//	    return fmt.Sprintf("UNALLOCATED(%032b)", i)
//	}
func (c Category) emitStringMethod(file *jen.File, s Schema) {
	recv := c.recv()
	formats := c.Encoding.Formats

	file.Func().
		Params(c.recvParam()).
		Id("String").Params().String().
		BlockFunc(func(g *jen.Group) {
			g.Id("mnemonic").Op(":=").Add(recv).Dot("Mnemonic").Call()

			locals := neededLocals(formats)
			if len(locals) > 0 {
				g.Line()
				for _, name := range locals {
					f := s.mustField(name)
					g.Id(name).Op(":=").Add(f.getCall(recv))
				}
				g.Line()
			}

			for _, format := range formats {
				sprintfArgs := make([]jen.Code, 0, 1+len(format.Arguments))
				sprintfArgs = append(sprintfArgs, jen.Lit(format.Format))
				for _, arg := range format.Arguments {
					sprintfArgs = append(sprintfArgs, jen.Id(arg))
				}
				ret := jen.Return(jen.Qual("fmt", "Sprintf").Call(sprintfArgs...))

				if format.Condition == "" {
					g.Add(ret)
					continue
				}

				g.If(parseCondition(format.Condition)).Block(ret)
				g.Line()
			}
		})

	file.Line()
}

// collapseParams collapses consecutive same-type fields into a single parameter
// group: (rd, rn, rm Register) instead of (rd Register, rn Register, rm Register)
func collapseParams(fields []SchemaField) []jen.Code {
	var params []jen.Code
	for i := 0; i < len(fields); {
		lastIndex := i + 1
		for lastIndex < len(fields) && fields[lastIndex].Type == fields[i].Type {
			lastIndex++
		}

		var idents []jen.Code
		for _, f := range fields[i:lastIndex] {
			idents = append(idents, jen.Id(f.Name))
		}

		params = append(params, jen.List(idents...).Add(fields[i].Type.jenCode()))
		i = lastIndex
	}

	return params
}

// Spec is the top-level JSON spec object
type Spec struct {
	Arch       string     `json:"arch"`
	Package    string     `json:"package"`
	Categories []Category `json:"categories"`
}

// Generate produces one jennifer File per category, keyed by output filename
func (s Spec) Generate() (map[string]*jen.File, error) {
	files := make(map[string]*jen.File, len(s.Categories))
	for _, cat := range s.Categories {
		f, err := cat.generate(s.Package)
		if err != nil {
			return nil, fmt.Errorf("category %s: %w", cat.Name, err)
		}
		files[camelToSnake(cat.Name)+".go"] = f
	}
	return files, nil
}

func main() {
	specPath := flag.String("spec", "", "path to the JSON spec file (required)")
	outDir := flag.String("out", ".", "output directory")
	flag.Parse()

	if *specPath == "" {
		log.Fatal("--spec is required")
	}

	data, err := os.ReadFile(*specPath)
	if err != nil {
		log.Fatalf("read spec: %v", err)
	}

	var spec Spec
	if err := json.Unmarshal(data, &spec); err != nil {
		log.Fatalf("parse spec: %v", err)
	}

	if err := os.MkdirAll(*outDir, 0o755); err != nil {
		log.Fatalf("create output directory: %v", err)
	}

	files, err := spec.Generate()
	if err != nil {
		log.Fatal(err)
	}

	for name, file := range files {
		path := filepath.Join(*outDir, name)

		f, err := os.Create(path)
		if err != nil {
			log.Fatalf("create %s: %v", path, err)
		}

		if renderErr := file.Render(f); renderErr != nil {
			f.Close()
			log.Fatalf("render %s: %v", path, renderErr)
		}

		f.Close()
		fmt.Printf("generated: %s\n", path)
	}
}
