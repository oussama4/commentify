// package validate provides some reusable validation helpers
package validate

type ValidationError map[string]string

func (e ValidationError) Error() string {
	return "invalid input"
}

func (e ValidationError) Fields() map[string]string {
	return e
}

type Validator struct {
	Err ValidationError
}

func New() *Validator {
	return &Validator{Err: make(ValidationError)}
}

func (v *Validator) Valid() error {
	if len(v.Err) != 0 {
		return v.Err
	}
	return nil
}

func (v *Validator) Check(valid bool, field string, message string) {
	if !valid {
		v.Err[field] = message
	}
}
