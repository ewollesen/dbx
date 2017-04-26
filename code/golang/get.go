// Copyright (C) 2017 Space Monkey, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package golang

import (
	"strings"

	"bitbucket.org/pkg/inflect"
	"gopkg.in/spacemonkeygo/dbx.v1/ir"
	"gopkg.in/spacemonkeygo/dbx.v1/sql"
	"gopkg.in/spacemonkeygo/dbx.v1/sqlgen/sqlembedgo"
)

type Get struct {
	PartitionedArgs
	Info   sqlembedgo.Info
	Suffix string
	Row    *Var
	LastPk *Var
}

func GetFromIR(ir_read *ir.Read, dialect sql.Dialect) *Get {
	select_sql := sql.SelectSQL(ir_read, dialect)
	get := &Get{
		PartitionedArgs: PartitionedArgsFromWheres(ir_read.Where),
		Info:            sqlembedgo.Embed("__", select_sql),
		Suffix:          convertSuffix(ir_read.Suffix),
	}

	get.Row = GetRowFromIR(ir_read)

	if ir_read.View == ir.Paged {
		pk_var := VarFromSelectable(ir_read.From.BasicPrimaryKey())
		pk_var.Name = "__" + pk_var.Name
		get.LastPk = pk_var
	}

	return get
}

func GetRowFromIR(ir_read *ir.Read) *Var {
	if model := ir_read.SelectedModel(); model != nil {
		return VarFromSelectable(model)
	}

	s := ResultStructFromRead(ir_read)
	return StructVar("row", s.Name, s.FieldVars())
}

func ResultStructFromRead(ir_read *ir.Read) *Struct {
	fields := FieldsFromSelectables(ir_read.Selectables)

	var parts []string
	for _, field := range fields {
		parts = append(parts, inflect.Camelize(field.Name))
	}
	parts = append(parts, "Row")

	return &Struct{
		Name:   strings.Join(parts, "_"),
		Fields: fields,
	}
}
