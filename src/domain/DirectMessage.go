package domain

var idDm int

// DirectMessage - Te maneja los mensajes papa
type DirectMessage struct {
	ID      int
	User    string
	UserTo  string
	Message string
	Read    bool
}

// CreateMessage - Crea un nuevo tweet ingresando el texto y el usuario
func CreateMessage(usr string, usrto string, txt string) *DirectMessage {
	directM := DirectMessage{
		generateIDDm(),
		usr,
		usrto,
		txt,
		false,
	}
	return &directM
}

// generateID - Genera un ID a cada nuevo tweet
func generateIDDm() int {
	idDm++
	return idDm
}
