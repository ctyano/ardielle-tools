//
// This file generated by rdl 1.4.11
//

package swagger

import (
	rdl "github.com/ardielle/ardielle-go/rdl"
)

var schema *rdl.Schema

func init() {
	sb := rdl.NewSchemaBuilder("swagger")
	sb.Comment("the swagger schema spec, expressed in RDL")

	tContact := rdl.NewStructTypeBuilder("Struct", "Contact")
	tContact.Field("name", "String", true, nil, "")
	tContact.Field("url", "String", true, nil, "")
	tContact.Field("email", "String", true, nil, "")
	sb.AddType(tContact.Build())

	tLicense := rdl.NewStructTypeBuilder("Struct", "License")
	tLicense.Field("name", "String", true, nil, "")
	tLicense.Field("url", "String", true, nil, "")
	sb.AddType(tLicense.Build())

	tInfo := rdl.NewStructTypeBuilder("Struct", "Info")
	tInfo.Field("title", "String", false, nil, "")
	tInfo.Field("version", "String", true, nil, "")
	tInfo.Field("description", "String", true, nil, "")
	tInfo.Field("termsOfService", "String", true, nil, "")
	tInfo.Field("contact", "Contact", true, nil, "")
	tInfo.Field("license", "License", true, nil, "")
	sb.AddType(tInfo.Build())

	tType := rdl.NewMapTypeBuilder("Map", "Type")
	tType.Keys("String")
	tType.Items("Any")
	sb.AddType(tType.Build())

	tParameter := rdl.NewStructTypeBuilder("Struct", "Parameter")
	tParameter.Field("name", "String", false, nil, "")
	tParameter.Field("in", "String", false, nil, "\"query\", \"header\", \"path\", \"formData\", \"body\"")
	tParameter.Field("schema", "Type", true, nil, "")
	tParameter.Field("type", "String", true, nil, "")
	tParameter.Field("format", "String", true, nil, "")
	tParameter.Field("collectionFormat", "String", false, "csv", "")
	tParameter.Field("required", "Bool", false, false, "must be true for path params")
	tParameter.Field("description", "String", true, nil, "")
	sb.AddType(tParameter.Build())

	tResponse := rdl.NewStructTypeBuilder("Struct", "Response")
	tResponse.Field("description", "String", false, nil, "")
	tResponse.Field("schema", "Type", false, nil, "")
	sb.AddType(tResponse.Build())

	tOperation := rdl.NewStructTypeBuilder("Struct", "Operation")
	tOperation.ArrayField("tags", "String", true, "")
	tOperation.Field("summary", "String", true, nil, "")
	tOperation.Field("operationID", "String", true, nil, "")
	tOperation.ArrayField("consumes", "String", true, "")
	tOperation.ArrayField("produces", "String", true, "")
	tOperation.ArrayField("parameters", "Parameter", true, "")
	tOperation.MapField("responses", "String", "Response", false, "")
	sb.AddType(tOperation.Build())

	tPathItem := rdl.NewStructTypeBuilder("Struct", "PathItem")
	tPathItem.Field("ref", "String", true, nil, "")
	tPathItem.Field("get", "Operation", true, nil, "")
	tPathItem.Field("put", "Operation", true, nil, "")
	tPathItem.Field("post", "Operation", true, nil, "")
	tPathItem.Field("delete", "Operation", true, nil, "")
	tPathItem.Field("options", "Operation", true, nil, "")
	tPathItem.Field("head", "Operation", true, nil, "")
	tPathItem.Field("patch", "Operation", true, nil, "")
	sb.AddType(tPathItem.Build())

	tSecurityDef := rdl.NewStructTypeBuilder("Struct", "SecurityDef")
	tSecurityDef.Field("in", "String", false, nil, "")
	tSecurityDef.Field("name", "String", false, nil, "")
	tSecurityDef.Field("type", "String", false, nil, "")
	sb.AddType(tSecurityDef.Build())

	tDoc := rdl.NewStructTypeBuilder("Struct", "Doc")
	tDoc.Field("swagger", "String", false, nil, "always \"2.0\" for now")
	tDoc.Field("info", "Info", false, nil, "")
	tDoc.Field("basePath", "String", false, nil, "")
	tDoc.Field("host", "String", true, nil, "")
	tDoc.ArrayField("schemes", "String", true, "")
	tDoc.MapField("paths", "String", "PathItem", false, "")
	tDoc.MapField("definitions", "String", "Type", false, "")
	tDoc.MapField("securityDefinitions", "String", "SecurityDef", true, "")
	sb.AddType(tDoc.Build())

	schema = sb.Build()
}

func SwaggerSchema() *rdl.Schema {
	return schema
}
