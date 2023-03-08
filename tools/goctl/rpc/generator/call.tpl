{{.head}}

package {{.filePackage}}

import (
	{{.pbPackage}}
	{{if ne .pbPackage .protoGoPackage}}{{.protoGoPackage}}{{end}}
)

type (
	{{.alias}}
)



var NewClient = {{.pkg}}.New{{.serviceName}}Client