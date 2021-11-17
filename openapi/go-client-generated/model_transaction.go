/*
 * Access API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Transaction struct {
	Id string `json:"id"`
	Script string `json:"script"`
	Arguments []string `json:"arguments"`
	ReferenceBlockId string `json:"reference_block_id"`
	GasLimit int32 `json:"gas_limit"`
	Payer string `json:"payer"`
	ProposalKey *ProposalKey `json:"proposal_key,omitempty"`
	Authorizers []string `json:"authorizers,omitempty"`
	PayloadSignatures []TransactionSignature `json:"payload_signatures,omitempty"`
	EnvelopeSignatures []TransactionSignature `json:"envelope_signatures,omitempty"`
	Result *TransactionResult `json:"result,omitempty"`
	Expandable *TransactionExpandable `json:"_expandable,omitempty"`
	Links *Links `json:"_links,omitempty"`
}
