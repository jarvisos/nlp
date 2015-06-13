package network

// Network Interface
// This interface is used by Jarvis NLP to actually run any
// NLP problems. It contains the actual Neural Network to use
// for determining meaning.
type Network interface {
	// initalize Function
	// Used to run any setup required for the network
	//
	// return bool
	// 		a boolean value based on the success of the
	// 		setup
	Initalize(bool) bool

	// setInput Function
	// Used to set the input string for the network to
	// process
	//
	// string parameter
	//		A string for the network to process
	SetInput(string)

	// destroy Function
	// Used to run any internal shutdown the network needs
	// before it is released
	//
	// return bool
	// 		A boolean value based on the success of the
	//		destruction
	Destroy() bool
}
