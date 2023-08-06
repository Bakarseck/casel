package casel

type Error struct {
	MsgError string
}

var Errors = map[int]string{
	404: "La page que vous cherchez n'existe pas",
	500: "Erreur interne du serveur",
	403: "Accès interdit",
	405: "Méthode non autorisée",
}
