// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package otherpackage

import fmt "fmt"

func (dst *Embed) SetFields(src *Embed, paths ...string) error {
	if len(paths) != 0 {
		return fmt.Errorf("message Embed has no fields, but paths %s were specified", paths)
	}
	if src != nil {
		*dst = *src
	}
	return nil
}
