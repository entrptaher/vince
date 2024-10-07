// DO NOT EDIT Code generateVendord by ua/os/make_os.go
package ua2

var vendorAllRe = MatchRe(`(?:^|[^A-Z0-9-_]|[^A-Z0-9-]_|sprd-)(?:MAAR(JS)?||MAAUNP0[26789]ASJBASU2(JS)?||CPDTDFCPNTDF(JS?)CMNTDF(JS)?CMDTDF(JS)?||MDDR(JS)?MDDC(JS)?MDDS(JS)?||MAFS(JS)?FSJB||MAGW(JS)?||HPCMHPHPNTDF(JS)?HPDTDF(JS)?||MANM(JS)?||MALC(JS)?MALE(JS)?MALN(JS)?LCJBLEN2||MAMI(JS)?MAM3||MAMD||Ordissimowebissimo3||MASM(JS)?SMJB||MASE(JS)?MASP(JS)?MASA(JS)?||MATM(JS)?MATB(JS)?MATP(JS)?TNJBTAJB)`)
var vendorAll = []*vendorRe{
	{re: MatchRe(`(?:^|[^A-Z0-9-_]|[^A-Z0-9-]_|sprd-)(?:MAAR(JS)?)`), name: "Acer"},
	{re: MatchRe(`(?:^|[^A-Z0-9-_]|[^A-Z0-9-]_|sprd-)(?:MAAU|NP0[26789]|ASJB|ASU2(JS)?)`), name: "Asus"},
	{re: MatchRe(`(?:^|[^A-Z0-9-_]|[^A-Z0-9-]_|sprd-)(?:CPDTDF|CPNTDF(JS?)|CMNTDF(JS)?|CMDTDF(JS)?)`), name: "Compaq"},
	{re: MatchRe(`(?:^|[^A-Z0-9-_]|[^A-Z0-9-]_|sprd-)(?:MDDR(JS)?|MDDC(JS)?|MDDS(JS)?)`), name: "Dell"},
	{re: MatchRe(`(?:^|[^A-Z0-9-_]|[^A-Z0-9-]_|sprd-)(?:MAFS(JS)?|FSJB)`), name: "Fujitsu"},
	{re: MatchRe(`(?:^|[^A-Z0-9-_]|[^A-Z0-9-]_|sprd-)(?:MAGW(JS)?)`), name: "Gateway"},
	{re: MatchRe(`(?:^|[^A-Z0-9-_]|[^A-Z0-9-]_|sprd-)(?:HPCMHP|HPNTDF(JS)?|HPDTDF(JS)?)`), name: "HP"},
	{re: MatchRe(`(?:^|[^A-Z0-9-_]|[^A-Z0-9-]_|sprd-)(?:MANM(JS)?)`), name: "Hyrican"},
	{re: MatchRe(`(?:^|[^A-Z0-9-_]|[^A-Z0-9-]_|sprd-)(?:MALC(JS)?|MALE(JS)?|MALN(JS)?|LCJB|LEN2)`), name: "Lenovo"},
	{re: MatchRe(`(?:^|[^A-Z0-9-_]|[^A-Z0-9-]_|sprd-)(?:MAMI(JS)?|MAM3)`), name: "MSI"},
	{re: MatchRe(`(?:^|[^A-Z0-9-_]|[^A-Z0-9-]_|sprd-)(?:MAMD)`), name: "Medion"},
	{re: MatchRe(`(?:^|[^A-Z0-9-_]|[^A-Z0-9-]_|sprd-)(?:Ordissimo|webissimo3)`), name: "Ordissimo"},
	{re: MatchRe(`(?:^|[^A-Z0-9-_]|[^A-Z0-9-]_|sprd-)(?:MASM(JS)?|SMJB)`), name: "Samsung"},
	{re: MatchRe(`(?:^|[^A-Z0-9-_]|[^A-Z0-9-]_|sprd-)(?:MASE(JS)?|MASP(JS)?|MASA(JS)?)`), name: "Sony"},
	{re: MatchRe(`(?:^|[^A-Z0-9-_]|[^A-Z0-9-]_|sprd-)(?:MATM(JS)?|MATB(JS)?|MATP(JS)?|TNJB|TAJB)`), name: "Toshiba"},
}
