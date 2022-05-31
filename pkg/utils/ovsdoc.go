package utils

import "encoding/xml"

type Database struct {
	XMLName xml.Name `xml:"database"`
	Text    string   `xml:",chardata"`
	Name    string   `xml:"name,attr"`
	Title   string   `xml:"title,attr"`
	P       []struct {
		Text string `xml:",chardata"`
		Code string `xml:"code"`
		Ref  struct {
			Text string `xml:",chardata"`
			Db   string `xml:"db,attr"`
		} `xml:"ref"`
	} `xml:"p"`
	H2 string `xml:"h2"`
	Dl struct {
		Text string `xml:",chardata"`
		Dt   struct {
			Text string `xml:",chardata"`
			Code string `xml:"code"`
		} `xml:"dt"`
		Dd string `xml:"dd"`
	} `xml:"dl"`
	Table []struct {
		Text  string `xml:",chardata"`
		Name  string `xml:"name,attr"`
		Title string `xml:"title,attr"`
		P     []struct {
			Text string `xml:",chardata"`
			Ref  []struct {
				Text   string `xml:",chardata"`
				Table  string `xml:"table,attr"`
				Column string `xml:"column,attr"`
				Db     string `xml:"db,attr"`
			} `xml:"ref"`
			Code []string `xml:"code"`
		} `xml:"p"`
		Group []struct {
			Text   string `xml:",chardata"`
			Title  string `xml:"title,attr"`
			Column []struct {
				Text string   `xml:",chardata"`
				Name string   `xml:"name,attr"`
				Key  string   `xml:"key,attr"`
				Type string   `xml:"type,attr"`
				Code []string `xml:"code"`
				P    []struct {
					Text string `xml:",chardata"`
					Code []struct {
						Text string   `xml:",chardata"`
						Var  []string `xml:"var"`
					} `xml:"code"`
					Ref []struct {
						Text   string `xml:",chardata"`
						Column string `xml:"column,attr"`
						Key    string `xml:"key,attr"`
						Table  string `xml:"table,attr"`
						Db     string `xml:"db,attr"`
					} `xml:"ref"`
					Em  string `xml:"em"`
					Var string `xml:"var"`
					Ul  struct {
						Text string `xml:",chardata"`
						Li   []struct {
							Text string `xml:",chardata"`
							Code string `xml:"code"`
						} `xml:"li"`
					} `xml:"ul"`
				} `xml:"p"`
				Pre string `xml:"pre"`
				Ref []struct {
					Text   string `xml:",chardata"`
					Column string `xml:"column,attr"`
					Db     string `xml:"db,attr"`
					Table  string `xml:"table,attr"`
					Key    string `xml:"key,attr"`
				} `xml:"ref"`
				Em string `xml:"em"`
				Ul struct {
					Text string `xml:",chardata"`
					Li   []struct {
						Text string   `xml:",chardata"`
						Code string   `xml:"code"`
						P    []string `xml:"p"`
						Ref  []struct {
							Text   string `xml:",chardata"`
							Table  string `xml:"table,attr"`
							Column string `xml:"column,attr"`
						} `xml:"ref"`
					} `xml:"li"`
				} `xml:"ul"`
				Dl struct {
					Text string `xml:",chardata"`
					Dt   []struct {
						Text string `xml:",chardata"`
						Code []struct {
							Text string `xml:",chardata"`
							Var  string `xml:"var"`
						} `xml:"code"`
						Var []string `xml:"var"`
					} `xml:"dt"`
					Dd []struct {
						Text string `xml:",chardata"`
						Ref  []struct {
							Text   string `xml:",chardata"`
							Column string `xml:"column,attr"`
							Key    string `xml:"key,attr"`
							Table  string `xml:"table,attr"`
						} `xml:"ref"`
						Code []string `xml:"code"`
						P    []struct {
							Text string   `xml:",chardata"`
							Code []string `xml:"code"`
							Ref  []struct {
								Text   string `xml:",chardata"`
								Column string `xml:"column,attr"`
								Table  string `xml:"table,attr"`
								Db     string `xml:"db,attr"`
								Key    string `xml:"key,attr"`
							} `xml:"ref"`
							Var []string `xml:"var"`
						} `xml:"p"`
						Ul struct {
							Text string `xml:",chardata"`
							Li   []struct {
								Text string   `xml:",chardata"`
								Code []string `xml:"code"`
							} `xml:"li"`
						} `xml:"ul"`
						Dl struct {
							Text string `xml:",chardata"`
							Dt   []struct {
								Text string `xml:",chardata"`
								Code string `xml:"code"`
							} `xml:"dt"`
							Dd []string `xml:"dd"`
						} `xml:"dl"`
					} `xml:"dd"`
				} `xml:"dl"`
				Ol struct {
					Text string   `xml:",chardata"`
					Li   []string `xml:"li"`
				} `xml:"ol"`
			} `xml:"column"`
			Group []struct {
				Text  string `xml:",chardata"`
				Title string `xml:"title,attr"`
				P     []struct {
					Text string   `xml:",chardata"`
					Code []string `xml:"code"`
					Ref  []struct {
						Text   string `xml:",chardata"`
						Db     string `xml:"db,attr"`
						Column string `xml:"column,attr"`
						Table  string `xml:"table,attr"`
					} `xml:"ref"`
				} `xml:"p"`
				Column []struct {
					Text string   `xml:",chardata"`
					Name string   `xml:"name,attr"`
					Key  string   `xml:"key,attr"`
					Type string   `xml:"type,attr"`
					Code []string `xml:"code"`
					Ref  []struct {
						Text   string `xml:",chardata"`
						Db     string `xml:"db,attr"`
						Column string `xml:"column,attr"`
						Table  string `xml:"table,attr"`
						Key    string `xml:"key,attr"`
					} `xml:"ref"`
					P []struct {
						Text string   `xml:",chardata"`
						Code []string `xml:"code"`
						Ref  []struct {
							Text   string `xml:",chardata"`
							Column string `xml:"column,attr"`
							Table  string `xml:"table,attr"`
							Db     string `xml:"db,attr"`
						} `xml:"ref"`
					} `xml:"p"`
					Dl struct {
						Text string `xml:",chardata"`
						Dt   []struct {
							Text string `xml:",chardata"`
							Code string `xml:"code"`
						} `xml:"dt"`
						Dd []struct {
							Text string `xml:",chardata"`
							P    []struct {
								Text string `xml:",chardata"`
								Ref  []struct {
									Text   string `xml:",chardata"`
									Column string `xml:"column,attr"`
									Key    string `xml:"key,attr"`
								} `xml:"ref"`
								Code string `xml:"code"`
							} `xml:"p"`
						} `xml:"dd"`
					} `xml:"dl"`
					Ul struct {
						Text string `xml:",chardata"`
						Li   []struct {
							Text string `xml:",chardata"`
							Code string `xml:"code"`
						} `xml:"li"`
					} `xml:"ul"`
				} `xml:"column"`
				Group struct {
					Text   string `xml:",chardata"`
					Title  string `xml:"title,attr"`
					Column []struct {
						Text string `xml:",chardata"`
						Name string `xml:"name,attr"`
						Key  string `xml:"key,attr"`
						Ref  []struct {
							Text   string `xml:",chardata"`
							Table  string `xml:"table,attr"`
							Column string `xml:"column,attr"`
							Db     string `xml:"db,attr"`
						} `xml:"ref"`
					} `xml:"column"`
				} `xml:"group"`
			} `xml:"group"`
			P []struct {
				Text string `xml:",chardata"`
				Ref  []struct {
					Text   string `xml:",chardata"`
					Column string `xml:"column,attr"`
					Key    string `xml:"key,attr"`
					Table  string `xml:"table,attr"`
				} `xml:"ref"`
				Code []struct {
					Text string   `xml:",chardata"`
					Var  []string `xml:"var"`
				} `xml:"code"`
				Dfn string `xml:"dfn"`
			} `xml:"p"`
			Ul struct {
				Text string `xml:",chardata"`
				Li   []struct {
					Text string `xml:",chardata"`
					Ref  []struct {
						Text   string `xml:",chardata"`
						Table  string `xml:"table,attr"`
						Column string `xml:"column,attr"`
					} `xml:"ref"`
					Code []string `xml:"code"`
				} `xml:"li"`
			} `xml:"ul"`
			Code string `xml:"code"`
		} `xml:"group"`
		Column []struct {
			Text string `xml:",chardata"`
			Name string `xml:"name,attr"`
			Key  string `xml:"key,attr"`
			Em   string `xml:"em"`
			P    []struct {
				Text string `xml:",chardata"`
				Ref  []struct {
					Text   string `xml:",chardata"`
					Table  string `xml:"table,attr"`
					Column string `xml:"column,attr"`
					Db     string `xml:"db,attr"`
					Colun  string `xml:"colun,attr"`
					Key    string `xml:"key,attr"`
				} `xml:"ref"`
				Code []string `xml:"code"`
				Var  []string `xml:"var"`
				B    string   `xml:"b"`
			} `xml:"p"`
			Ref []struct {
				Text   string `xml:",chardata"`
				Table  string `xml:"table,attr"`
				Db     string `xml:"db,attr"`
				Column string `xml:"column,attr"`
			} `xml:"ref"`
			Code []string `xml:"code"`
			Ul   struct {
				Text string `xml:",chardata"`
				Li   []struct {
					Text string   `xml:",chardata"`
					Code []string `xml:"code"`
					Ref  []struct {
						Text   string `xml:",chardata"`
						Column string `xml:"column,attr"`
					} `xml:"ref"`
				} `xml:"li"`
			} `xml:"ul"`
		} `xml:"column"`
		Pre string `xml:"pre"`
	} `xml:"table"`
} 
