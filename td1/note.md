pkg /toto - toto1.go - func Test1 {}  // fonction globale visible après import (package "toto")
          - toto2.go - func test2 {}  // fonction privée (seulement entre les toto)
          - toto3.go

go mod init /Users/shubq/workflow/ia04-td/td1 // Specifier les versions utilisées, sert a savoir ou est la racine
