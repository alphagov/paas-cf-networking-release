package api_test

import (
	"policy-server/api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EgressDestinationsValidator", func() {
	var validator api.EgressDestinationsValidator

	BeforeEach(func() {
		validator = api.EgressDestinationsValidator{}
	})

	FDescribe("ValidateEgressDestinations", func() {
		It("does not error for valid egress destinations", func() {
			destinations := []api.EgressDestination{
				{
					Name:        "meow",
					Description: "a cat",
					Protocol:    "tcp",
					Ports:       []api.Ports{{Start: 8080, End: 8081}},
					IPRanges:    []api.IPRange{{Start: "10.265.0.0", End: "10.265.0.0"}},
				},
			}

			err := validator.ValidateEgressDestinations(destinations)
			Expect(err).NotTo(HaveOccurred())
		})

		Context("when the name is missing", func() {
			It("returns an error", func() {
				destinations := []api.EgressDestination{
					{
						Description: "a cat",
						Protocol:    "tcp",
						Ports:       []api.Ports{{Start: 8080, End: 8081}},
						IPRanges:    []api.IPRange{{Start: "10.265.0.0", End: "10.265.0.0"}},
					},
				}

				err := validator.ValidateEgressDestinations(destinations)
				Expect(err).To(MatchError("missing destination name"))
			})

		})

		Context("when no ips are provided", func() {
			It("returns an error", func() {
				destinations := []api.EgressDestination{
					{
						Name:        "meow",
						Description: "a cat",
						Protocol:    "tcp",
						Ports:       []api.Ports{{Start: 8080, End: 8081}},
						IPRanges:    []api.IPRange{},
					},
				}

				err := validator.ValidateEgressDestinations(destinations)
				Expect(err).To(MatchError("missing destination IP range"))
			})
		})

		Context("when no ports are provided", func() {
			It("returns an error", func() {
				destinations := []api.EgressDestination{
					{
						Name:        "meow",
						Description: "a cat",
						Protocol:    "tcp",
						Ports:       []api.Ports{},
						IPRanges:    []api.IPRange{{Start: "10.265.0.0", End: "10.265.0.0"}},
					},
				}

				err := validator.ValidateEgressDestinations(destinations)
				Expect(err).To(MatchError("missing destination ports"))
			})
		})

		Context("when no protocol is provided", func() {
			It("returns an error", func() {
				destinations := []api.EgressDestination{
					{
						Name:        "meow",
						Description: "a cat",
						Ports:       []api.Ports{{Start: 8080, End: 8081}},
						IPRanges:    []api.IPRange{{Start: "10.265.0.0", End: "10.265.0.0"}},
					},
				}

				err := validator.ValidateEgressDestinations(destinations)
				Expect(err).To(MatchError("missing destination protocol"))
			})
		})

		Context("when the protocol is not a supported type", func() {
			It("returns an error", func() {
				destinations := []api.EgressDestination{
					{
						Name:        "meow",
						Description: "a cat",
						Protocol:    "meow",
						Ports:       []api.Ports{{Start: 8080, End: 8081}},
						IPRanges:    []api.IPRange{{Start: "10.265.0.0", End: "10.265.0.0"}},
					},
				}

				err := validator.ValidateEgressDestinations(destinations)
				Expect(err).To(MatchError("invalid destination protocol 'meow', specify either tcp, udp, or icmp"))
			})
		})

		Context("when the protocol is not icmp", func() {
			Context("when the icmp type is provided", func() {
				It("returns an error", func() {
					It("returns an error", func() {
						icmpType := 13
						destinations := []api.EgressDestination{
							{
								Name:        "meow",
								Description: "a cat",
								Protocol:    "tcp",
								Ports:       []api.Ports{{Start: 8080, End: 8081}},
								IPRanges:    []api.IPRange{{Start: "10.265.0.0", End: "10.265.0.0"}},
								ICMPType:    &icmpType,
							},
						}

						err := validator.ValidateEgressDestinations(destinations)
						Expect(err).To(MatchError("invalid destination: cannot add icmp_type property destination with type 'tcp'"))
					})
				})
			})

			Context("when the icmp code is provided", func() {
				It("returns an error", func() {
					It("returns an error", func() {
						icmpCode := 13
						destinations := []api.EgressDestination{
							{
								Name:        "meow",
								Description: "a cat",
								Protocol:    "udp",
								Ports:       []api.Ports{{Start: 8080, End: 8081}},
								IPRanges:    []api.IPRange{{Start: "10.265.0.0", End: "10.265.0.0"}},
								ICMPCode:    &icmpCode,
							},
						}

						err := validator.ValidateEgressDestinations(destinations)
						Expect(err).To(MatchError("invalid destination: cannot add icmp_code property destination with type 'udp'"))
					})
				})
			})
		})

		//Context("when the policies list is nil", func() {
		//	It("returns a useful error", func() {
		//		err := validator.ValidatePolicies(nil)
		//		Expect(err).To(MatchError("missing policies"))
		//	})
		//})
		//
		//Context("when the policies list is empty", func() {
		//	It("returns a useful error", func() {
		//		err := validator.ValidatePolicies([]api.Policy{})
		//		Expect(err).To(MatchError("missing policies"))
		//	})
		//})
		//
		//Context("when source id is missing", func() {
		//	It("returns a useful error", func() {
		//		policies := []api.Policy{
		//			api.Policy{
		//				Source: api.Source{
		//					ID:  "",
		//					Tag: "",
		//				},
		//				Destination: api.Destination{
		//					ID:       "some-destination-id",
		//					Tag:      "",
		//					Protocol: "tcp",
		//					Ports: api.Ports{
		//						Start: 42,
		//						End:   42,
		//					},
		//				},
		//			},
		//		}
		//
		//		err := validator.ValidatePolicies(policies)
		//		Expect(err).To(MatchError("missing source id"))
		//	})
		//})
		//
		//Context("when destination id is missing", func() {
		//	It("returns a useful error", func() {
		//		policies := []api.Policy{
		//			api.Policy{
		//				Source: api.Source{
		//					ID:  "foo",
		//					Tag: "",
		//				},
		//				Destination: api.Destination{
		//					ID:       "",
		//					Tag:      "",
		//					Protocol: "tcp",
		//					Ports: api.Ports{
		//						Start: 42,
		//						End:   42,
		//					},
		//				},
		//			},
		//		}
		//
		//		err := validator.ValidatePolicies(policies)
		//		Expect(err).To(MatchError("missing destination id"))
		//	})
		//})
		//
		//Context("when invalid destination protocol", func() {
		//	It("returns a useful error", func() {
		//		policies := []api.Policy{
		//			api.Policy{
		//				Source: api.Source{
		//					ID:  "foo",
		//					Tag: "",
		//				},
		//				Destination: api.Destination{
		//					ID:       "bar",
		//					Tag:      "",
		//					Protocol: "banana",
		//					Ports: api.Ports{
		//						Start: 42,
		//						End:   42,
		//					},
		//				},
		//			},
		//		}
		//
		//		err := validator.ValidatePolicies(policies)
		//		Expect(err).To(MatchError("invalid destination protocol, specify either udp or tcp"))
		//	})
		//})
		//
		//Context("when the end port is less than the start port", func() {
		//	It("returns a useful error", func() {
		//		policies := []api.Policy{
		//			api.Policy{
		//				Source: api.Source{
		//					ID:  "foo",
		//					Tag: "",
		//				},
		//				Destination: api.Destination{
		//					ID:       "bar",
		//					Tag:      "",
		//					Protocol: "tcp",
		//					Ports: api.Ports{
		//						Start: 1243,
		//						End:   999,
		//					},
		//				},
		//			},
		//		}
		//
		//		err := validator.ValidatePolicies(policies)
		//		Expect(err).To(MatchError("invalid port range 1243-999, start must be less than or equal to end"))
		//	})
		//})
		//
		//Context("when the start port is less than 0", func() {
		//	It("returns a useful error", func() {
		//		policies := []api.Policy{
		//			api.Policy{
		//				Source: api.Source{
		//					ID:  "foo",
		//					Tag: "",
		//				},
		//				Destination: api.Destination{
		//					ID:       "bar",
		//					Tag:      "",
		//					Protocol: "tcp",
		//					Ports: api.Ports{
		//						Start: -42,
		//						End:   999,
		//					},
		//				},
		//			},
		//		}
		//
		//		err := validator.ValidatePolicies(policies)
		//		Expect(err).To(MatchError("invalid start port -42, must be in range 1-65535"))
		//	})
		//})
		//
		//Context("when the start port is missing", func() {
		//	It("returns a useful error", func() {
		//		policies := []api.Policy{
		//			api.Policy{
		//				Source: api.Source{
		//					ID:  "foo",
		//					Tag: "",
		//				},
		//				Destination: api.Destination{
		//					ID:       "bar",
		//					Tag:      "",
		//					Protocol: "tcp",
		//					Ports: api.Ports{
		//						Start: 0,
		//						End:   999,
		//					},
		//				},
		//			},
		//		}
		//
		//		err := validator.ValidatePolicies(policies)
		//		Expect(err).To(MatchError("missing start port"))
		//	})
		//})
		//
		//Context("when the end port is greater than 65535", func() {
		//	It("returns a useful error", func() {
		//		policies := []api.Policy{
		//			api.Policy{
		//				Source: api.Source{
		//					ID:  "foo",
		//					Tag: "",
		//				},
		//				Destination: api.Destination{
		//					ID:       "bar",
		//					Tag:      "",
		//					Protocol: "tcp",
		//					Ports: api.Ports{
		//						Start: 42,
		//						End:   101101,
		//					},
		//				},
		//			},
		//		}
		//
		//		err := validator.ValidatePolicies(policies)
		//		Expect(err).To(MatchError("invalid end port 101101, must be in range 1-65535"))
		//	})
		//})
		//
		//Context("when a tag is supplied", func() {
		//	It("returns a useful error", func() {
		//		policies := []api.Policy{
		//			{
		//				Source: api.Source{
		//					ID:  "foo",
		//					Tag: "some-tag",
		//				},
		//				Destination: api.Destination{
		//					ID:       "bar",
		//					Tag:      "",
		//					Protocol: "tcp",
		//					Ports: api.Ports{
		//						Start: 123,
		//						End:   456,
		//					},
		//				},
		//			},
		//		}
		//
		//		err := validator.ValidatePolicies(policies)
		//		Expect(err).To(MatchError("tags may not be specified"))
		//	})
		//})
	})
})
