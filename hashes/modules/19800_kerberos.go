package modules

import (
	"github.com/tankbusta/hashvalidate/hashes"
	"github.com/tankbusta/hashvalidate/tokenizer"
)

type kerberos_19800 struct{}

func init() {
	hashes.Register(19800, kerberos_19800{})
}

func (s kerberos_19800) Name() string { return "Kerberos 5, etype 17, Pre-Auth" }

func (s kerberos_19800) Example() string {
	return "$krb5pa$17$hashcat$HASHCATDOMAIN.COM$a17776abe5383236c58582f515843e029ecbff43706d177651b7b6cdb2713b17597ddb35b1c9c470c281589fd1d51cca125414d19e40e333"
}

func (s kerberos_19800) Type() int { return 19800 }

func (s kerberos_19800) Tokens() []tokenizer.Token {
	// The rest of this is generated by GenerateTokens based on the incoming hash
	return []tokenizer.Token{
		{
			Length:     11,
			Attributes: tokenizer.FixedLength | tokenizer.VerifySignature,
			Signatures: []tokenizer.Signature{
				{
					Expected: "$krb5pa$17$",
				},
			},
		},
		{
			Separator:  "$",
			LengthMin:  1,
			LengthMax:  512,
			Attributes: tokenizer.VerifyLength,
		},
		{
			Separator:  "$",
			LengthMin:  1,
			LengthMax:  512,
			Attributes: tokenizer.VerifyLength,
		},
		{
			Separator:  "$",
			LengthMin:  104,
			LengthMax:  112,
			Attributes: tokenizer.VerifyLength | tokenizer.VerifyHex,
		},
	}
}
