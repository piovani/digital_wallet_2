package operation

import "fmt"

func ValidInput(input Input) error {
	err := ValidStringEmpty(input.UserName)
	if err != nil {
		return fmt.Errorf("user_name: %s", err.Error())
	}

	err = ValidStringEmpty(input.Coin)
	if err != nil {
		return fmt.Errorf("coin: %s", err.Error())
	}

	err = ValidValueFloat(input.Value)
	if err != nil {
		return fmt.Errorf("value: %s", err.Error())
	}

	return nil
}

func ValidStringEmpty(s string) error {
	if s == "" {
		return fmt.Errorf("must be informed")
	}

	return nil
}

func ValidValueFloat(f float64) error {
	if f <= 0 {
		return fmt.Errorf("must be greater than 0")
	}

	return nil
}
