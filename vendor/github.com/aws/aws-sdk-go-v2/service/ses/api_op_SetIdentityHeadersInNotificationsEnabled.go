// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Given an identity (an email address or a domain), sets whether Amazon SES
// includes the original email headers in the Amazon Simple Notification Service
// (Amazon SNS) notifications of a specified type.
//
// You can execute this operation no more than once per second.
//
// For more information about using notifications with Amazon SES, see the [Amazon SES Developer Guide].
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity-using-notifications.html
func (c *Client) SetIdentityHeadersInNotificationsEnabled(ctx context.Context, params *SetIdentityHeadersInNotificationsEnabledInput, optFns ...func(*Options)) (*SetIdentityHeadersInNotificationsEnabledOutput, error) {
	if params == nil {
		params = &SetIdentityHeadersInNotificationsEnabledInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetIdentityHeadersInNotificationsEnabled", params, optFns, c.addOperationSetIdentityHeadersInNotificationsEnabledMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetIdentityHeadersInNotificationsEnabledOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to set whether Amazon SES includes the original email
// headers in the Amazon SNS notifications of a specified type. For information
// about notifications, see the [Amazon SES Developer Guide].
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity-using-notifications-sns.html
type SetIdentityHeadersInNotificationsEnabledInput struct {

	// Sets whether Amazon SES includes the original email headers in Amazon SNS
	// notifications of the specified notification type. A value of true specifies
	// that Amazon SES includes headers in notifications, and a value of false
	// specifies that Amazon SES does not include headers in notifications.
	//
	// This value can only be set when NotificationType is already set to use a
	// particular Amazon SNS topic.
	//
	// This member is required.
	Enabled bool

	// The identity for which to enable or disable headers in notifications. Examples:
	// user@example.com , example.com .
	//
	// This member is required.
	Identity *string

	// The notification type for which to enable or disable headers in notifications.
	//
	// This member is required.
	NotificationType types.NotificationType

	noSmithyDocumentSerde
}

// An empty element returned on a successful request.
type SetIdentityHeadersInNotificationsEnabledOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetIdentityHeadersInNotificationsEnabledMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetIdentityHeadersInNotificationsEnabled{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetIdentityHeadersInNotificationsEnabled{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SetIdentityHeadersInNotificationsEnabled"); err != nil {
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
	if err = addOpSetIdentityHeadersInNotificationsEnabledValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetIdentityHeadersInNotificationsEnabled(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetIdentityHeadersInNotificationsEnabled(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SetIdentityHeadersInNotificationsEnabled",
	}
}