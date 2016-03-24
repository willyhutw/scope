// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package wafiface provides an interface for the AWS WAF.
package wafiface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/waf"
)

// WAFAPI is the interface type for waf.WAF.
type WAFAPI interface {
	CreateByteMatchSetRequest(*waf.CreateByteMatchSetInput) (*request.Request, *waf.CreateByteMatchSetOutput)

	CreateByteMatchSet(*waf.CreateByteMatchSetInput) (*waf.CreateByteMatchSetOutput, error)

	CreateIPSetRequest(*waf.CreateIPSetInput) (*request.Request, *waf.CreateIPSetOutput)

	CreateIPSet(*waf.CreateIPSetInput) (*waf.CreateIPSetOutput, error)

	CreateRuleRequest(*waf.CreateRuleInput) (*request.Request, *waf.CreateRuleOutput)

	CreateRule(*waf.CreateRuleInput) (*waf.CreateRuleOutput, error)

	CreateSqlInjectionMatchSetRequest(*waf.CreateSqlInjectionMatchSetInput) (*request.Request, *waf.CreateSqlInjectionMatchSetOutput)

	CreateSqlInjectionMatchSet(*waf.CreateSqlInjectionMatchSetInput) (*waf.CreateSqlInjectionMatchSetOutput, error)

	CreateWebACLRequest(*waf.CreateWebACLInput) (*request.Request, *waf.CreateWebACLOutput)

	CreateWebACL(*waf.CreateWebACLInput) (*waf.CreateWebACLOutput, error)

	DeleteByteMatchSetRequest(*waf.DeleteByteMatchSetInput) (*request.Request, *waf.DeleteByteMatchSetOutput)

	DeleteByteMatchSet(*waf.DeleteByteMatchSetInput) (*waf.DeleteByteMatchSetOutput, error)

	DeleteIPSetRequest(*waf.DeleteIPSetInput) (*request.Request, *waf.DeleteIPSetOutput)

	DeleteIPSet(*waf.DeleteIPSetInput) (*waf.DeleteIPSetOutput, error)

	DeleteRuleRequest(*waf.DeleteRuleInput) (*request.Request, *waf.DeleteRuleOutput)

	DeleteRule(*waf.DeleteRuleInput) (*waf.DeleteRuleOutput, error)

	DeleteSqlInjectionMatchSetRequest(*waf.DeleteSqlInjectionMatchSetInput) (*request.Request, *waf.DeleteSqlInjectionMatchSetOutput)

	DeleteSqlInjectionMatchSet(*waf.DeleteSqlInjectionMatchSetInput) (*waf.DeleteSqlInjectionMatchSetOutput, error)

	DeleteWebACLRequest(*waf.DeleteWebACLInput) (*request.Request, *waf.DeleteWebACLOutput)

	DeleteWebACL(*waf.DeleteWebACLInput) (*waf.DeleteWebACLOutput, error)

	GetByteMatchSetRequest(*waf.GetByteMatchSetInput) (*request.Request, *waf.GetByteMatchSetOutput)

	GetByteMatchSet(*waf.GetByteMatchSetInput) (*waf.GetByteMatchSetOutput, error)

	GetChangeTokenRequest(*waf.GetChangeTokenInput) (*request.Request, *waf.GetChangeTokenOutput)

	GetChangeToken(*waf.GetChangeTokenInput) (*waf.GetChangeTokenOutput, error)

	GetChangeTokenStatusRequest(*waf.GetChangeTokenStatusInput) (*request.Request, *waf.GetChangeTokenStatusOutput)

	GetChangeTokenStatus(*waf.GetChangeTokenStatusInput) (*waf.GetChangeTokenStatusOutput, error)

	GetIPSetRequest(*waf.GetIPSetInput) (*request.Request, *waf.GetIPSetOutput)

	GetIPSet(*waf.GetIPSetInput) (*waf.GetIPSetOutput, error)

	GetRuleRequest(*waf.GetRuleInput) (*request.Request, *waf.GetRuleOutput)

	GetRule(*waf.GetRuleInput) (*waf.GetRuleOutput, error)

	GetSampledRequestsRequest(*waf.GetSampledRequestsInput) (*request.Request, *waf.GetSampledRequestsOutput)

	GetSampledRequests(*waf.GetSampledRequestsInput) (*waf.GetSampledRequestsOutput, error)

	GetSqlInjectionMatchSetRequest(*waf.GetSqlInjectionMatchSetInput) (*request.Request, *waf.GetSqlInjectionMatchSetOutput)

	GetSqlInjectionMatchSet(*waf.GetSqlInjectionMatchSetInput) (*waf.GetSqlInjectionMatchSetOutput, error)

	GetWebACLRequest(*waf.GetWebACLInput) (*request.Request, *waf.GetWebACLOutput)

	GetWebACL(*waf.GetWebACLInput) (*waf.GetWebACLOutput, error)

	ListByteMatchSetsRequest(*waf.ListByteMatchSetsInput) (*request.Request, *waf.ListByteMatchSetsOutput)

	ListByteMatchSets(*waf.ListByteMatchSetsInput) (*waf.ListByteMatchSetsOutput, error)

	ListIPSetsRequest(*waf.ListIPSetsInput) (*request.Request, *waf.ListIPSetsOutput)

	ListIPSets(*waf.ListIPSetsInput) (*waf.ListIPSetsOutput, error)

	ListRulesRequest(*waf.ListRulesInput) (*request.Request, *waf.ListRulesOutput)

	ListRules(*waf.ListRulesInput) (*waf.ListRulesOutput, error)

	ListSqlInjectionMatchSetsRequest(*waf.ListSqlInjectionMatchSetsInput) (*request.Request, *waf.ListSqlInjectionMatchSetsOutput)

	ListSqlInjectionMatchSets(*waf.ListSqlInjectionMatchSetsInput) (*waf.ListSqlInjectionMatchSetsOutput, error)

	ListWebACLsRequest(*waf.ListWebACLsInput) (*request.Request, *waf.ListWebACLsOutput)

	ListWebACLs(*waf.ListWebACLsInput) (*waf.ListWebACLsOutput, error)

	UpdateByteMatchSetRequest(*waf.UpdateByteMatchSetInput) (*request.Request, *waf.UpdateByteMatchSetOutput)

	UpdateByteMatchSet(*waf.UpdateByteMatchSetInput) (*waf.UpdateByteMatchSetOutput, error)

	UpdateIPSetRequest(*waf.UpdateIPSetInput) (*request.Request, *waf.UpdateIPSetOutput)

	UpdateIPSet(*waf.UpdateIPSetInput) (*waf.UpdateIPSetOutput, error)

	UpdateRuleRequest(*waf.UpdateRuleInput) (*request.Request, *waf.UpdateRuleOutput)

	UpdateRule(*waf.UpdateRuleInput) (*waf.UpdateRuleOutput, error)

	UpdateSqlInjectionMatchSetRequest(*waf.UpdateSqlInjectionMatchSetInput) (*request.Request, *waf.UpdateSqlInjectionMatchSetOutput)

	UpdateSqlInjectionMatchSet(*waf.UpdateSqlInjectionMatchSetInput) (*waf.UpdateSqlInjectionMatchSetOutput, error)

	UpdateWebACLRequest(*waf.UpdateWebACLInput) (*request.Request, *waf.UpdateWebACLOutput)

	UpdateWebACL(*waf.UpdateWebACLInput) (*waf.UpdateWebACLOutput, error)
}

var _ WAFAPI = (*waf.WAF)(nil)
