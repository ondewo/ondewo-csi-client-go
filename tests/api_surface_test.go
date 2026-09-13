// Copyright 2020-2026 ONDEWO GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Everything in this file names the ONDEWO CSI API specifically: its services, its messages, its
// enums. It is the ONLY product-specific file of the suite - generated_code_test.go and
// auth_test.go are product agnostic.
//
// CSI is a composite product: ondewo-csi-api vendors the NLU, S2T and T2S protos, so this module
// ships FOUR go packages and exposes 19 services - its own ondewo.csi.Conversations plus the
// sixteen ondewo.nlu.*, ondewo.s2t.Speech2Text and ondewo.t2s.Text2Speech. A consumer that dials a
// CSI server and then talks NLU to it over the same client is the normal case, so all four
// packages are part of this client's surface and all four are pinned here.
package tests

import (
	"context"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"

	csi "github.com/ondewo/ondewo-csi-client-go/v5/api/ondewo/csi"
	nlu "github.com/ondewo/ondewo-csi-client-go/v5/api/ondewo/nlu"
	s2t "github.com/ondewo/ondewo-csi-client-go/v5/api/ondewo/s2t"
	t2s "github.com/ondewo/ondewo-csi-client-go/v5/api/ondewo/t2s"
)

// protoFileCount is the number of .proto files below ondewo-csi-api/ondewo that the compiler
// consumed. Every one of them has to end up in the global descriptor registry when this package
// is linked; a proto that silently stopped being compiled is otherwise invisible until a
// consumer misses a type.
const protoFileCount = 21

// services is every gRPC service this product exposes, keyed by the fully qualified proto name
// the ServiceDesc must declare.
var services = map[string]*grpc.ServiceDesc{
	"ondewo.csi.Conversations":     &csi.Conversations_ServiceDesc,
	"ondewo.nlu.Agents":            &nlu.Agents_ServiceDesc,
	"ondewo.nlu.AiServices":        &nlu.AiServices_ServiceDesc,
	"ondewo.nlu.CcaiProjects":      &nlu.CcaiProjects_ServiceDesc,
	"ondewo.nlu.Contexts":          &nlu.Contexts_ServiceDesc,
	"ondewo.nlu.EntityTypes":       &nlu.EntityTypes_ServiceDesc,
	"ondewo.nlu.Intents":           &nlu.Intents_ServiceDesc,
	"ondewo.nlu.LlmEvaluations":    &nlu.LlmEvaluations_ServiceDesc,
	"ondewo.nlu.Operations":        &nlu.Operations_ServiceDesc,
	"ondewo.nlu.ProjectRoles":      &nlu.ProjectRoles_ServiceDesc,
	"ondewo.nlu.ProjectStatistics": &nlu.ProjectStatistics_ServiceDesc,
	"ondewo.nlu.Rags":              &nlu.Rags_ServiceDesc,
	"ondewo.nlu.ServerStatistics":  &nlu.ServerStatistics_ServiceDesc,
	"ondewo.nlu.Sessions":          &nlu.Sessions_ServiceDesc,
	"ondewo.nlu.Users":             &nlu.Users_ServiceDesc,
	"ondewo.nlu.Utilities":         &nlu.Utilities_ServiceDesc,
	"ondewo.nlu.Webhook":           &nlu.Webhook_ServiceDesc,
	"ondewo.s2t.Speech2Text":       &s2t.Speech2Text_ServiceDesc,
	"ondewo.t2s.Text2Speech":       &t2s.Text2Speech_ServiceDesc,
}

// clientConstructors is the generated New<Service>Client of every service above. A client SDK
// that compiles but whose constructors are missing is useless, and the two generators that
// produce them (protoc-gen-go, protoc-gen-go-grpc) can disagree - so both halves are listed.
var clientConstructors = map[string]func(grpc.ClientConnInterface) any{
	"ondewo.csi.Conversations":     func(cc grpc.ClientConnInterface) any { return csi.NewConversationsClient(cc) },
	"ondewo.nlu.Agents":            func(cc grpc.ClientConnInterface) any { return nlu.NewAgentsClient(cc) },
	"ondewo.nlu.AiServices":        func(cc grpc.ClientConnInterface) any { return nlu.NewAiServicesClient(cc) },
	"ondewo.nlu.CcaiProjects":      func(cc grpc.ClientConnInterface) any { return nlu.NewCcaiProjectsClient(cc) },
	"ondewo.nlu.Contexts":          func(cc grpc.ClientConnInterface) any { return nlu.NewContextsClient(cc) },
	"ondewo.nlu.EntityTypes":       func(cc grpc.ClientConnInterface) any { return nlu.NewEntityTypesClient(cc) },
	"ondewo.nlu.Intents":           func(cc grpc.ClientConnInterface) any { return nlu.NewIntentsClient(cc) },
	"ondewo.nlu.LlmEvaluations":    func(cc grpc.ClientConnInterface) any { return nlu.NewLlmEvaluationsClient(cc) },
	"ondewo.nlu.Operations":        func(cc grpc.ClientConnInterface) any { return nlu.NewOperationsClient(cc) },
	"ondewo.nlu.ProjectRoles":      func(cc grpc.ClientConnInterface) any { return nlu.NewProjectRolesClient(cc) },
	"ondewo.nlu.ProjectStatistics": func(cc grpc.ClientConnInterface) any { return nlu.NewProjectStatisticsClient(cc) },
	"ondewo.nlu.Rags":              func(cc grpc.ClientConnInterface) any { return nlu.NewRagsClient(cc) },
	"ondewo.nlu.ServerStatistics":  func(cc grpc.ClientConnInterface) any { return nlu.NewServerStatisticsClient(cc) },
	"ondewo.nlu.Sessions":          func(cc grpc.ClientConnInterface) any { return nlu.NewSessionsClient(cc) },
	"ondewo.nlu.Users":             func(cc grpc.ClientConnInterface) any { return nlu.NewUsersClient(cc) },
	"ondewo.nlu.Utilities":         func(cc grpc.ClientConnInterface) any { return nlu.NewUtilitiesClient(cc) },
	"ondewo.nlu.Webhook":           func(cc grpc.ClientConnInterface) any { return nlu.NewWebhookClient(cc) },
	"ondewo.s2t.Speech2Text":       func(cc grpc.ClientConnInterface) any { return s2t.NewSpeech2TextClient(cc) },
	"ondewo.t2s.Text2Speech":       func(cc grpc.ClientConnInterface) any { return t2s.NewText2SpeechClient(cc) },
}

// expectedMethods pins RPCs by name. The descriptor cross-check in generated_code_test.go proves
// the two generators agree with each other; it cannot notice an RPC that was renamed upstream,
// because both halves would be renamed together. These are spelled out so that a rename is a
// failing test rather than a silently broken consumer. The names are the ones ON THE WIRE (the
// `MethodName` of the ServiceDesc), which differ in case from the go identifiers wherever protoc
// re-capitalised an initialism - `GetS2sPipeline` on the wire is `GetS2SPipeline` in go.
var expectedMethods = map[string][]string{
	// S2sStream and GetControlStream are bidirectional/server streaming, so they live in
	// ServiceDesc.Streams rather than .Methods - the lookup has to consider both.
	"ondewo.csi.Conversations": {
		"CreateS2sPipeline", "GetS2sPipeline", "UpdateS2sPipeline", "DeleteS2sPipeline",
		"ListS2sPipelines", "CheckUpstreamHealth", "SetControlStatus",
		"S2sStream", "GetControlStream",
	},
	"ondewo.s2t.Speech2Text": {"TranscribeFile", "TranscribeStream", "GetS2tPipeline", "ListS2tPipelines"},
	"ondewo.t2s.Text2Speech": {"Synthesize", "BatchSynthesize", "StreamingSynthesize", "ListT2sPipelines"},
	"ondewo.nlu.Sessions":    {"DetectIntent", "StreamingDetectIntent", "ListSessions", "GetSession"},
	"ondewo.nlu.Agents":      {"CreateAgent", "GetAgent", "UpdateAgent", "DeleteAgent", "ListAgents"},
}

// TestMessageRoundTripsThroughTheWire is the core assertion about generated message code: a value
// built in go, serialized and parsed back is the same value. It covers scalars of three widths, a
// oneof carrying a nested message and a well-known Struct, so a generator that mis-numbers a field
// or loses a nested type fails here.
func TestMessageRoundTripsThroughTheWire(t *testing.T) {
	t.Parallel()

	content, err := structpb.NewStruct(map[string]any{"transfer_to": "+4312345678", "priority": float64(2)})
	if err != nil {
		t.Fatalf("structpb.NewStruct failed: %v", err)
	}

	original := &csi.S2SStreamResponse{
		Response: &csi.S2SStreamResponse_SipTrigger{
			SipTrigger: &csi.SipTrigger{Type: csi.SipTrigger_TRANSFER, Content: content},
		},
		UtteranceId: "4c9a1f2e-0f3a-4a53-9a4c-1b0a5a6f7c8d",
		ChunkIndex:  3,
		LastChunk:   true,
		TurnEpoch:   1 << 40,
	}

	wire, err := proto.Marshal(original)
	if err != nil {
		t.Fatalf("proto.Marshal(%T) failed: %v", original, err)
	}
	if len(wire) == 0 {
		t.Fatal("proto.Marshal produced 0 bytes for a fully populated message")
	}

	parsed := &csi.S2SStreamResponse{}
	if err := proto.Unmarshal(wire, parsed); err != nil {
		t.Fatalf("proto.Unmarshal failed: %v", err)
	}

	if !proto.Equal(original, parsed) {
		t.Fatalf("round trip changed the message:\n original = %v\n  parsed = %v", original, parsed)
	}
	if got, want := parsed.GetSipTrigger().GetType(), csi.SipTrigger_TRANSFER; got != want {
		t.Errorf("oneof branch after round trip = %v, want %v", got, want)
	}
	if got, want := parsed.GetSipTrigger().GetContent().GetFields()["transfer_to"].GetStringValue(), "+4312345678"; got != want {
		t.Errorf("well-known Struct value after round trip = %q, want %q", got, want)
	}
	if got, want := parsed.GetTurnEpoch(), uint64(1<<40); got != want {
		t.Errorf("uint64 after round trip = %d, want %d", got, want)
	}
}

// TestProto3ExplicitPresenceSurvivesTheWire guards the field kind that generators get wrong: a
// proto3 `optional` scalar has to keep the difference between "set to the zero value" and "not
// set". The angular target of the same compiler lost exactly this distinction, which made a
// false/0/"" unsendable; protoc-gen-go models it as a pointer, and this asserts it stays that way.
//
// ondewo/csi/conversation.proto itself declares no `optional` field; the one used here comes from
// the S2T protos ondewo-csi-api vendors, which this module compiles and publishes just the same.
func TestProto3ExplicitPresenceSurvivesTheWire(t *testing.T) {
	t.Parallel()

	t.Run("zero value set explicitly is transmitted", func(t *testing.T) {
		t.Parallel()

		wire, err := proto.Marshal(&s2t.S2TCloudProviderConfigAmazon{
			EnablePartialResultsStabilization: proto.Bool(false),
		})
		if err != nil {
			t.Fatalf("proto.Marshal failed: %v", err)
		}
		if len(wire) == 0 {
			t.Fatal("an explicitly set zero value was dropped from the wire - proto3 presence is lost")
		}

		parsed := &s2t.S2TCloudProviderConfigAmazon{}
		if err := proto.Unmarshal(wire, parsed); err != nil {
			t.Fatalf("proto.Unmarshal failed: %v", err)
		}
		if parsed.EnablePartialResultsStabilization == nil {
			t.Fatal("EnablePartialResultsStabilization is nil after the round trip, want a pointer to false")
		}
		if got := *parsed.EnablePartialResultsStabilization; got {
			t.Errorf("EnablePartialResultsStabilization = %v, want false", got)
		}
	})

	t.Run("unset stays unset", func(t *testing.T) {
		t.Parallel()

		wire, err := proto.Marshal(&s2t.S2TCloudProviderConfigAmazon{
			VocabularyName: proto.String("unset"),
		})
		if err != nil {
			t.Fatalf("proto.Marshal failed: %v", err)
		}

		parsed := &s2t.S2TCloudProviderConfigAmazon{}
		if err := proto.Unmarshal(wire, parsed); err != nil {
			t.Fatalf("proto.Unmarshal failed: %v", err)
		}
		if parsed.EnablePartialResultsStabilization != nil {
			t.Errorf(
				"EnablePartialResultsStabilization = %v after a round trip that never set it, want nil",
				*parsed.EnablePartialResultsStabilization,
			)
		}
	})
}

// TestEnumZeroValueIsTheUnspecifiedMember checks the member every proto3 enum should have at 0 and
// the name maps generated beside it. A zero value that is a real choice rather than "unspecified"
// is unrequestable in several of the other clients of this API - and ondewo.csi.ControlStatus is
// exactly such an enum, so its zero member is pinned by name instead of being wished away.
func TestEnumZeroValueIsTheUnspecifiedMember(t *testing.T) {
	t.Parallel()

	var zero csi.SipTrigger_SipTriggerType

	if zero != csi.SipTrigger_UNSPECIFIED {
		t.Errorf("zero value of SipTrigger_SipTriggerType = %v, want UNSPECIFIED", zero)
	}
	if got, want := zero.String(), "UNSPECIFIED"; got != want {
		t.Errorf("SipTrigger_SipTriggerType(0).String() = %q, want %q", got, want)
	}
	if got, want := csi.SipTrigger_SipTriggerType_name[0], "UNSPECIFIED"; got != want {
		t.Errorf("SipTrigger_SipTriggerType_name[0] = %q, want %q", got, want)
	}
	if got, want := csi.SipTrigger_SipTriggerType_value["HANGUP"], int32(csi.SipTrigger_HANGUP); got != want {
		t.Errorf("SipTrigger_SipTriggerType_value[HANGUP] = %d, want %d", got, want)
	}
	if got, want := int32(csi.SipTrigger_TRANSFER), int32(5); got != want {
		t.Errorf("TRANSFER = %d, want %d", got, want)
	}

	// ControlStatus has no *_UNSPECIFIED member: 0 is the live state OK, which the control stream
	// really sends. Pinned so a member inserted at 0 upstream - which would silently re-label every
	// stored zero - fails the build instead.
	var control csi.ControlStatus
	if control != csi.ControlStatus_OK {
		t.Errorf("zero value of ControlStatus = %v, want OK", control)
	}
	if got, want := csi.ControlStatus_name[0], "OK"; got != want {
		t.Errorf("ControlStatus_name[0] = %q, want %q", got, want)
	}
}

// TestUnmarshalRejectsTruncatedInput asserts the generated message reports a parse error instead
// of accepting a malformed payload: field 1 (`id`) is announced as 5 bytes long but only 1 follows.
func TestUnmarshalRejectsTruncatedInput(t *testing.T) {
	t.Parallel()

	if err := proto.Unmarshal([]byte{0x0a, 0x05, 'a'}, &csi.S2SPipeline{}); err == nil {
		t.Fatal("proto.Unmarshal accepted a truncated payload, want an error")
	}
}

// conversationsServer is a fake ONDEWO CSI server: it answers GetS2sPipeline and inherits the
// "unimplemented" behaviour of the generated base type for every other RPC of the service.
type conversationsServer struct {
	csi.UnimplementedConversationsServer
}

func (conversationsServer) GetS2SPipeline(_ context.Context, req *csi.S2SPipelineId) (*csi.S2SPipeline, error) {
	return &csi.S2SPipeline{
		Id:              req.GetId(),
		S2TPipelineId:   "default_german",
		NluProjectId:    "ae33586b-x2s2-494a-aa73-1af0589cfc56",
		NluLanguageCode: "de",
		T2SPipelineId:   "kerstin",
	}, nil
}

// TestUnaryRPCRoundTripsOverAnInProcessServer drives the generated client stub, the generated
// server stub and the generated ServiceDesc against each other over a real gRPC connection - the
// request is marshalled, routed by the method name baked into the stub, and the response is
// parsed back. Nothing here is mocked except the transport, which is in memory.
func TestUnaryRPCRoundTripsOverAnInProcessServer(t *testing.T) {
	t.Parallel()

	conn := dialInProcess(t, nil, func(srv *grpc.Server) {
		csi.RegisterConversationsServer(srv, conversationsServer{})
	})
	client := csi.NewConversationsClient(conn)

	const pipelineID = "pizza"
	response, err := client.GetS2SPipeline(t.Context(), &csi.S2SPipelineId{Id: pipelineID})
	if err != nil {
		t.Fatalf("GetS2sPipeline failed: %v", err)
	}

	if got := response.GetId(); got != pipelineID {
		t.Errorf("response pipeline id = %q, want %q", got, pipelineID)
	}
	if got, want := response.GetNluLanguageCode(), "de"; got != want {
		t.Errorf("response nlu language code = %q, want %q", got, want)
	}
}

// TestUnimplementedMethodIsReportedAsUnimplemented pins the other half of the generated server
// contract: an RPC the server does not implement must come back as codes.Unimplemented, not as a
// routing failure or a panic. It also proves the method is routed at all - a method missing from
// the ServiceDesc would surface as a different code.
func TestUnimplementedMethodIsReportedAsUnimplemented(t *testing.T) {
	t.Parallel()

	conn := dialInProcess(t, nil, func(srv *grpc.Server) {
		csi.RegisterConversationsServer(srv, conversationsServer{})
	})
	client := csi.NewConversationsClient(conn)

	_, err := client.ListS2SPipelines(t.Context(), &csi.ListS2SPipelinesRequest{})
	if got := status.Code(err); got != codes.Unimplemented {
		t.Fatalf("ListS2sPipelines returned code %v (err = %v), want %v", got, err, codes.Unimplemented)
	}
}
