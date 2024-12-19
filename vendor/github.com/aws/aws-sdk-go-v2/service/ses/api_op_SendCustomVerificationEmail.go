// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds an email address to the list of identities for your Amazon SES account in
// the current Amazon Web Services Region and attempts to verify it. As a result of
// executing this operation, a customized verification email is sent to the
// specified address.
//
// To use this operation, you must first create a custom verification email
// template. For more information about creating and using custom verification
// email templates, see [Using Custom Verification Email Templates]in the Amazon SES Developer Guide.
//
// You can execute this operation no more than once per second.
//
// [Using Custom Verification Email Templates]: https://docs.aws.amazon.com/ses/latest/dg/creating-identities.html#send-email-verify-address-custom
func (c *Client) SendCustomVerificationEmail(ctx context.Context, params *SendCustomVerificationEmailInput, optFns ...func(*Options)) (*SendCustomVerificationEmailOutput, error) {
	if params == nil {
		params = &SendCustomVerificationEmailInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendCustomVerificationEmail", params, optFns, c.addOperationSendCustomVerificationEmailMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendCustomVerificationEmailOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to send a custom verification email to a specified
// recipient.
type SendCustomVerificationEmailInput struct {

	// The email address to verify.
	//
	// This member is required.
	EmailAddress *string

	// The name of the custom verification email template to use when sending the
	// verification email.
	//
	// This member is required.
	TemplateName *string

	// Name of a configuration set to use when sending the verification email.
	ConfigurationSetName *string

	noSmithyDocumentSerde
}

// The response received when attempting to send the custom verification email.
type SendCustomVerificationEmailOutput struct {

	// The unique message identifier returned from the SendCustomVerificationEmail
	// operation.
	MessageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendCustomVerificationEmailMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSendCustomVerificationEmail{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSendCustomVerificationEmail{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SendCustomVerificationEmail"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpSendCustomVerificationEmailValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendCustomVerificationEmail(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opSendCustomVerificationEmail(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SendCustomVerificationEmail",
	}
}