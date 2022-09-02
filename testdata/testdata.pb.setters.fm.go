// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package testdata

import fmt "fmt"

func (dst *Empty) SetFields(src *Empty, paths ...string) error {
	if len(paths) != 0 {
		return fmt.Errorf("message Empty has no fields, but paths %s were specified", paths)
	}
	return nil
}

func (dst *Test) SetFields(src *Test, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "a":
			if len(subs) > 0 {
				var newDst, newSrc *Test_TestNested
				if (src == nil || src.A == nil) && dst.A == nil {
					continue
				}
				if src != nil {
					newSrc = src.A
				}
				if dst.A != nil {
					newDst = dst.A
				} else {
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
				var newDst, newSrc *Test_TestNested
				if (src == nil || src.B == nil) && dst.B == nil {
					continue
				}
				if src != nil {
					newSrc = src.B
				}
				if dst.B != nil {
					newDst = dst.B
				} else {
					newDst = &Test_TestNested{}
					dst.B = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.B = src.B
				} else {
					dst.B = nil
				}
			}
		case "c":
			if len(subs) > 0 {
				var newDst, newSrc *Test_TestNested
				if (src == nil || src.C == nil) && dst.C == nil {
					continue
				}
				if src != nil {
					newSrc = src.C
				}
				if dst.C != nil {
					newDst = dst.C
				} else {
					newDst = &Test_TestNested{}
					dst.C = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.C = src.C
				} else {
					dst.C = nil
				}
			}
		case "g":
			if len(subs) > 0 {
				var newDst, newSrc *Empty
				if (src == nil || src.G == nil) && dst.G == nil {
					continue
				}
				if src != nil {
					newSrc = src.G
				}
				if dst.G != nil {
					newDst = dst.G
				} else {
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
				dst.I = nil
			}
		case "j":
			if len(subs) > 0 {
				return fmt.Errorf("'j' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.J = src.J
			} else {
				dst.J = nil
			}
		case "l":
			if len(subs) > 0 {
				return fmt.Errorf("'l' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.L = src.L
			} else {
				var zero string
				dst.L = zero
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
					var srcTypeOk bool
					if src != nil {
						_, srcTypeOk = src.TestOneof.(*Test_D)
					}
					if srcValid := srcTypeOk || src == nil || src.TestOneof == nil || len(oneofSubs) == 0; !srcValid {
						return fmt.Errorf("attempt to set oneof 'd', while different oneof is set in source")
					}
					_, dstTypeOk := dst.TestOneof.(*Test_D)
					if dstValid := dstTypeOk || dst.TestOneof == nil || len(oneofSubs) == 0; !dstValid {
						return fmt.Errorf("attempt to set oneof 'd', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						return fmt.Errorf("'d' has no subfields, but %s were specified", oneofSubs)
					}
					if srcTypeOk {
						dst.TestOneof = src.TestOneof
					} else {
						dst.TestOneof = nil
					}
				case "e":
					var srcTypeOk bool
					if src != nil {
						_, srcTypeOk = src.TestOneof.(*Test_E)
					}
					if srcValid := srcTypeOk || src == nil || src.TestOneof == nil || len(oneofSubs) == 0; !srcValid {
						return fmt.Errorf("attempt to set oneof 'e', while different oneof is set in source")
					}
					_, dstTypeOk := dst.TestOneof.(*Test_E)
					if dstValid := dstTypeOk || dst.TestOneof == nil || len(oneofSubs) == 0; !dstValid {
						return fmt.Errorf("attempt to set oneof 'e', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						return fmt.Errorf("'e' has no subfields, but %s were specified", oneofSubs)
					}
					if srcTypeOk {
						dst.TestOneof = src.TestOneof
					} else {
						dst.TestOneof = nil
					}
				case "f":
					var srcTypeOk bool
					if src != nil {
						_, srcTypeOk = src.TestOneof.(*Test_F)
					}
					if srcValid := srcTypeOk || src == nil || src.TestOneof == nil || len(oneofSubs) == 0; !srcValid {
						return fmt.Errorf("attempt to set oneof 'f', while different oneof is set in source")
					}
					_, dstTypeOk := dst.TestOneof.(*Test_F)
					if dstValid := dstTypeOk || dst.TestOneof == nil || len(oneofSubs) == 0; !dstValid {
						return fmt.Errorf("attempt to set oneof 'f', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						return fmt.Errorf("'f' has no subfields, but %s were specified", oneofSubs)
					}
					if srcTypeOk {
						dst.TestOneof = src.TestOneof
					} else {
						dst.TestOneof = nil
					}
				case "k":
					var srcTypeOk bool
					if src != nil {
						_, srcTypeOk = src.TestOneof.(*Test_K)
					}
					if srcValid := srcTypeOk || src == nil || src.TestOneof == nil || len(oneofSubs) == 0; !srcValid {
						return fmt.Errorf("attempt to set oneof 'k', while different oneof is set in source")
					}
					_, dstTypeOk := dst.TestOneof.(*Test_K)
					if dstValid := dstTypeOk || dst.TestOneof == nil || len(oneofSubs) == 0; !dstValid {
						return fmt.Errorf("attempt to set oneof 'k', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *Test_TestNested
						if srcTypeOk {
							newSrc = src.TestOneof.(*Test_K).K
						}
						if dstTypeOk {
							newDst = dst.TestOneof.(*Test_K).K
						} else if srcTypeOk {
							newDst = &Test_TestNested{}
							dst.TestOneof = &Test_K{K: newDst}
						} else {
							dst.TestOneof = nil
							continue
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if srcTypeOk {
							dst.TestOneof = src.TestOneof
						} else {
							dst.TestOneof = nil
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
	for name, subs := range _processPaths(paths) {
		switch name {
		case "a":
			if len(subs) > 0 {
				var newDst, newSrc *Test_TestNested_TestNestedNested
				if (src == nil || src.A == nil) && dst.A == nil {
					continue
				}
				if src != nil {
					newSrc = src.A
				}
				if dst.A != nil {
					newDst = dst.A
				} else {
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
				var zero string
				dst.E = zero
			}
		case "f":
			if len(subs) > 0 {
				return fmt.Errorf("'f' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.F = src.F
			} else {
				var zero string
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
	for name, subs := range _processPaths(paths) {
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
				var newDst, newSrc *Test_TestNested_TestNestedNested_TestNestedNestedEmbed
				if (src == nil || src.H == nil) && dst.H == nil {
					continue
				}
				if src != nil {
					newSrc = src.H
				}
				if dst.H != nil {
					newDst = dst.H
				} else {
					newDst = &Test_TestNested_TestNestedNested_TestNestedNestedEmbed{}
					dst.H = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.H = src.H
				} else {
					dst.H = nil
				}
			}
		case "i":
			if len(subs) > 0 {
				var newDst, newSrc *Test_TestNested_TestNestedNested_TestNestedNestedEmbed2
				if (src == nil || src.I == nil) && dst.I == nil {
					continue
				}
				if src != nil {
					newSrc = src.I
				}
				if dst.I != nil {
					newDst = dst.I
				} else {
					newDst = &Test_TestNested_TestNestedNested_TestNestedNestedEmbed2{}
					dst.I = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.I = src.I
				} else {
					dst.I = nil
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
					var srcTypeOk bool
					if src != nil {
						_, srcTypeOk = src.TestNestedNestedOneOf.(*Test_TestNested_TestNestedNested_E)
					}
					if srcValid := srcTypeOk || src == nil || src.TestNestedNestedOneOf == nil || len(oneofSubs) == 0; !srcValid {
						return fmt.Errorf("attempt to set oneof 'e', while different oneof is set in source")
					}
					_, dstTypeOk := dst.TestNestedNestedOneOf.(*Test_TestNested_TestNestedNested_E)
					if dstValid := dstTypeOk || dst.TestNestedNestedOneOf == nil || len(oneofSubs) == 0; !dstValid {
						return fmt.Errorf("attempt to set oneof 'e', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *Empty
						if srcTypeOk {
							newSrc = src.TestNestedNestedOneOf.(*Test_TestNested_TestNestedNested_E).E
						}
						if dstTypeOk {
							newDst = dst.TestNestedNestedOneOf.(*Test_TestNested_TestNestedNested_E).E
						} else if srcTypeOk {
							newDst = &Empty{}
							dst.TestNestedNestedOneOf = &Test_TestNested_TestNestedNested_E{E: newDst}
						} else {
							dst.TestNestedNestedOneOf = nil
							continue
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if srcTypeOk {
							dst.TestNestedNestedOneOf = src.TestNestedNestedOneOf
						} else {
							dst.TestNestedNestedOneOf = nil
						}
					}
				case "f":
					var srcTypeOk bool
					if src != nil {
						_, srcTypeOk = src.TestNestedNestedOneOf.(*Test_TestNested_TestNestedNested_F)
					}
					if srcValid := srcTypeOk || src == nil || src.TestNestedNestedOneOf == nil || len(oneofSubs) == 0; !srcValid {
						return fmt.Errorf("attempt to set oneof 'f', while different oneof is set in source")
					}
					_, dstTypeOk := dst.TestNestedNestedOneOf.(*Test_TestNested_TestNestedNested_F)
					if dstValid := dstTypeOk || dst.TestNestedNestedOneOf == nil || len(oneofSubs) == 0; !dstValid {
						return fmt.Errorf("attempt to set oneof 'f', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						return fmt.Errorf("'f' has no subfields, but %s were specified", oneofSubs)
					}
					if srcTypeOk {
						dst.TestNestedNestedOneOf = src.TestNestedNestedOneOf
					} else {
						dst.TestNestedNestedOneOf = nil
					}
				case "g":
					var srcTypeOk bool
					if src != nil {
						_, srcTypeOk = src.TestNestedNestedOneOf.(*Test_TestNested_TestNestedNested_G)
					}
					if srcValid := srcTypeOk || src == nil || src.TestNestedNestedOneOf == nil || len(oneofSubs) == 0; !srcValid {
						return fmt.Errorf("attempt to set oneof 'g', while different oneof is set in source")
					}
					_, dstTypeOk := dst.TestNestedNestedOneOf.(*Test_TestNested_TestNestedNested_G)
					if dstValid := dstTypeOk || dst.TestNestedNestedOneOf == nil || len(oneofSubs) == 0; !dstValid {
						return fmt.Errorf("attempt to set oneof 'g', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						return fmt.Errorf("'g' has no subfields, but %s were specified", oneofSubs)
					}
					if srcTypeOk {
						dst.TestNestedNestedOneOf = src.TestNestedNestedOneOf
					} else {
						dst.TestNestedNestedOneOf = nil
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
	for name, subs := range _processPaths(paths) {
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
	for name, subs := range _processPaths(paths) {
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
