// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package testdata

import (
	fmt "fmt"
	github_com_TheThingsIndustries_protoc_gen_fieldmask_testdata_testpackage "github.com/TheThingsIndustries/protoc-gen-fieldmask/testdata/testpackage"
	types "github.com/gogo/protobuf/types"
)

func (dst *Empty) SetFields(src *Empty, paths ...string) error {
	if len(paths) != 0 {
		return fmt.Errorf("message Empty has no fields, but paths %s were specified", paths)
	}
	if src != nil {
		*dst = *src
	}
	return nil
}

func (dst *Test) SetFields(src *Test, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "a":
			if len(subs) > 0 {
				var newSrc *Test_TestNested
				if src != nil {
					newSrc = src.A
				}
				newDst := dst.A
				if newDst == nil {
					if newSrc == nil {
						continue
					}
					newDst = &Test_TestNested{}
					dst.A = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.A = src.A
				} else {
					dst.A = nil
				}
			}
		case "b":
			if len(subs) > 0 {
				var newSrc *Test_TestNested
				if src != nil {
					newSrc = src.CustomName
				}
				newDst := dst.CustomName
				if newDst == nil {
					if newSrc == nil {
						continue
					}
					newDst = &Test_TestNested{}
					dst.CustomName = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.CustomName = src.CustomName
				} else {
					dst.CustomName = nil
				}
			}
		case "c":
			if len(subs) > 0 {
				newDst := &dst.C
				var newSrc *Test_TestNested
				if src != nil {
					newSrc = &src.C
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.C = src.C
				} else {
					var zero Test_TestNested
					dst.C = zero
				}
			}
		case "g":
			if len(subs) > 0 {
				var newSrc *Empty
				if src != nil {
					newSrc = src.G
				}
				newDst := dst.G
				if newDst == nil {
					if newSrc == nil {
						continue
					}
					newDst = &Empty{}
					dst.G = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.G = src.G
				} else {
					dst.G = nil
				}
			}
		case "h":
			if len(subs) > 0 {
				return fmt.Errorf("'h' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.H = src.H
			} else {
				dst.H = nil
			}
		case "i":
			if len(subs) > 0 {
				return fmt.Errorf("'i' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.I = src.I
			} else {
				var zero types.StringValue
				dst.I = zero
			}

		case "testOneof":
			if len(subs) == 0 && src == nil {
				dst.TestOneof = nil
				continue
			} else if len(subs) == 0 {
				dst.TestOneof = src.TestOneof
				continue
			}

			subPathMap := _processPaths(subs)
			if len(subPathMap) > 1 {
				return fmt.Errorf("more than one field specified for oneof field '%s'", name)
			}
			for oneofName, oneofSubs := range subPathMap {
				switch oneofName {
				case "d":
					if _, ok := dst.TestOneof.(*Test_D); !ok {
						dst.TestOneof = &Test_D{}
					}
					if len(oneofSubs) > 0 {
						return fmt.Errorf("'d' has no subfields, but %s were specified", oneofSubs)
					}
					if src != nil {
						dst.TestOneof.(*Test_D).D = src.GetD()
					} else {
						var zero int32
						dst.TestOneof.(*Test_D).D = zero
					}
				case "e":
					if _, ok := dst.TestOneof.(*Test_CustomNameOneof); !ok {
						dst.TestOneof = &Test_CustomNameOneof{}
					}
					if len(oneofSubs) > 0 {
						return fmt.Errorf("'e' has no subfields, but %s were specified", oneofSubs)
					}
					if src != nil {
						dst.TestOneof.(*Test_CustomNameOneof).CustomNameOneof = src.GetCustomNameOneof()
					} else {
						var zero uint32
						dst.TestOneof.(*Test_CustomNameOneof).CustomNameOneof = zero
					}
				case "f":
					if _, ok := dst.TestOneof.(*Test_F); !ok {
						dst.TestOneof = &Test_F{}
					}
					if len(oneofSubs) > 0 {
						return fmt.Errorf("'f' has no subfields, but %s were specified", oneofSubs)
					}
					if src != nil {
						dst.TestOneof.(*Test_F).F = src.GetF()
					} else {
						dst.TestOneof.(*Test_F).F = nil
					}
				case "k":
					if _, ok := dst.TestOneof.(*Test_K); !ok {
						dst.TestOneof = &Test_K{}
					}
					if len(oneofSubs) > 0 {
						var newSrc *Test_TestNested
						if src != nil {
							newSrc = src.GetK()
						}
						newDst := dst.TestOneof.(*Test_K).K
						if newDst == nil {
							if newSrc == nil {
								continue
							}
							newDst = &Test_TestNested{}
							dst.TestOneof.(*Test_K).K = newDst
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.TestOneof.(*Test_K).K = src.GetK()
						} else {
							dst.TestOneof.(*Test_K).K = nil
						}
					}

				default:
					return fmt.Errorf("invalid oneof field: '%s.%s'", name, oneofName)
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Test_TestNested) SetFields(src *Test_TestNested, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "a":
			if len(subs) > 0 {
				var newSrc *Test_TestNested_TestNestedNested
				if src != nil {
					newSrc = src.A
				}
				newDst := dst.A
				if newDst == nil {
					if newSrc == nil {
						continue
					}
					newDst = &Test_TestNested_TestNestedNested{}
					dst.A = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.A = src.A
				} else {
					dst.A = nil
				}
			}
		case "b":
			if len(subs) > 0 {
				return fmt.Errorf("'b' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.B = src.B
			} else {
				dst.B = nil
			}
		case "c":
			if len(subs) > 0 {
				return fmt.Errorf("'c' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.C = src.C
			} else {
				dst.C = nil
			}
		case "d":
			if len(subs) > 0 {
				return fmt.Errorf("'d' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.D = src.D
			} else {
				dst.D = nil
			}
		case "e":
			if len(subs) > 0 {
				return fmt.Errorf("'e' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.E = src.E
			} else {
				dst.E = nil
			}
		case "f":
			if len(subs) > 0 {
				return fmt.Errorf("'f' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.F = src.F
			} else {
				var zero github_com_TheThingsIndustries_protoc_gen_fieldmask_testdata_testpackage.CustomType
				dst.F = zero
			}
		case "g":
			if len(subs) > 0 {
				return fmt.Errorf("'g' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.G = src.G
			} else {
				dst.G = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Test_TestNested_TestNestedNested) SetFields(src *Test_TestNested_TestNestedNested, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "a":
			if len(subs) > 0 {
				return fmt.Errorf("'a' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.A = src.A
			} else {
				var zero int32
				dst.A = zero
			}
		case "b":
			if len(subs) > 0 {
				return fmt.Errorf("'b' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.B = src.B
			} else {
				var zero int64
				dst.B = zero
			}
		case "c":
			if len(subs) > 0 {
				return fmt.Errorf("'c' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.C = src.C
			} else {
				dst.C = nil
			}
		case "d":
			if len(subs) > 0 {
				return fmt.Errorf("'d' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.D = src.D
			} else {
				dst.D = nil
			}
		case "h":
			if len(subs) > 0 {
				var newSrc *Test_TestNested_TestNestedNested_TestNestedNestedEmbed
				if src != nil {
					newSrc = src.Test_TestNested_TestNestedNested_TestNestedNestedEmbed
				}
				newDst := dst.Test_TestNested_TestNestedNested_TestNestedNestedEmbed
				if newDst == nil {
					if newSrc == nil {
						continue
					}
					newDst = &Test_TestNested_TestNestedNested_TestNestedNestedEmbed{}
					dst.Test_TestNested_TestNestedNested_TestNestedNestedEmbed = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Test_TestNested_TestNestedNested_TestNestedNestedEmbed = src.Test_TestNested_TestNestedNested_TestNestedNestedEmbed
				} else {
					dst.Test_TestNested_TestNestedNested_TestNestedNestedEmbed = nil
				}
			}
		case "i":
			if len(subs) > 0 {
				newDst := &dst.Test_TestNested_TestNestedNested_TestNestedNestedEmbed2
				var newSrc *Test_TestNested_TestNestedNested_TestNestedNestedEmbed2
				if src != nil {
					newSrc = &src.Test_TestNested_TestNestedNested_TestNestedNestedEmbed2
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Test_TestNested_TestNestedNested_TestNestedNestedEmbed2 = src.Test_TestNested_TestNestedNested_TestNestedNestedEmbed2
				} else {
					var zero Test_TestNested_TestNestedNested_TestNestedNestedEmbed2
					dst.Test_TestNested_TestNestedNested_TestNestedNestedEmbed2 = zero
				}
			}

		case "testNestedNestedOneOf":
			if len(subs) == 0 && src == nil {
				dst.TestNestedNestedOneOf = nil
				continue
			} else if len(subs) == 0 {
				dst.TestNestedNestedOneOf = src.TestNestedNestedOneOf
				continue
			}

			subPathMap := _processPaths(subs)
			if len(subPathMap) > 1 {
				return fmt.Errorf("more than one field specified for oneof field '%s'", name)
			}
			for oneofName, oneofSubs := range subPathMap {
				switch oneofName {
				case "e":
					if _, ok := dst.TestNestedNestedOneOf.(*Test_TestNested_TestNestedNested_E); !ok {
						dst.TestNestedNestedOneOf = &Test_TestNested_TestNestedNested_E{}
					}
					if len(oneofSubs) > 0 {
						var newSrc *Empty
						if src != nil {
							newSrc = src.GetE()
						}
						newDst := dst.TestNestedNestedOneOf.(*Test_TestNested_TestNestedNested_E).E
						if newDst == nil {
							if newSrc == nil {
								continue
							}
							newDst = &Empty{}
							dst.TestNestedNestedOneOf.(*Test_TestNested_TestNestedNested_E).E = newDst
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.TestNestedNestedOneOf.(*Test_TestNested_TestNestedNested_E).E = src.GetE()
						} else {
							dst.TestNestedNestedOneOf.(*Test_TestNested_TestNestedNested_E).E = nil
						}
					}
				case "f":
					if _, ok := dst.TestNestedNestedOneOf.(*Test_TestNested_TestNestedNested_F); !ok {
						dst.TestNestedNestedOneOf = &Test_TestNested_TestNestedNested_F{}
					}
					if len(oneofSubs) > 0 {
						return fmt.Errorf("'f' has no subfields, but %s were specified", oneofSubs)
					}
					if src != nil {
						dst.TestNestedNestedOneOf.(*Test_TestNested_TestNestedNested_F).F = src.GetF()
					} else {
						var zero uint32
						dst.TestNestedNestedOneOf.(*Test_TestNested_TestNestedNested_F).F = zero
					}
				case "g":
					if _, ok := dst.TestNestedNestedOneOf.(*Test_TestNested_TestNestedNested_G); !ok {
						dst.TestNestedNestedOneOf = &Test_TestNested_TestNestedNested_G{}
					}
					if len(oneofSubs) > 0 {
						return fmt.Errorf("'g' has no subfields, but %s were specified", oneofSubs)
					}
					if src != nil {
						dst.TestNestedNestedOneOf.(*Test_TestNested_TestNestedNested_G).G = src.GetG()
					} else {
						dst.TestNestedNestedOneOf.(*Test_TestNested_TestNestedNested_G).G = nil
					}

				default:
					return fmt.Errorf("invalid oneof field: '%s.%s'", name, oneofName)
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Test_TestNested_TestNestedNested_TestNestedNestedEmbed) SetFields(src *Test_TestNested_TestNestedNested_TestNestedNestedEmbed, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "nested_field":
			if len(subs) > 0 {
				return fmt.Errorf("'nested_field' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.NestedField = src.NestedField
			} else {
				var zero int32
				dst.NestedField = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Test_TestNested_TestNestedNested_TestNestedNestedEmbed2) SetFields(src *Test_TestNested_TestNestedNested_TestNestedNestedEmbed2, paths ...string) error {
	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		switch name {
		case "nested_field_2":
			if len(subs) > 0 {
				return fmt.Errorf("'nested_field_2' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.NestedField_2 = src.NestedField_2
			} else {
				var zero int32
				dst.NestedField_2 = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
