package playground

func PrintAST(script string) error {
	parser := &parser{
		Buffer: script,
		Pretty: true,
	}

	if err := parser.Init(); err != nil {
		return err
	}
	if err := parser.Parse(); err != nil {
		return err
	}

	parser.PrintSyntaxTree()

	return nil
}
