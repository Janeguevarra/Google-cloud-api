// Copyright 2017 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pstest

import (
	"fmt"
	"io"
	"sync"
	"testing"
	"time"

	"cloud.google.com/go/internal/testutil"
	"github.com/golang/protobuf/ptypes"
	"golang.org/x/net/context"
	pb "google.golang.org/genproto/googleapis/pubsub/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestTopics(t *testing.T) {
	pclient, _, server := newFake(t)
	defer server.Close()

	ctx := context.Background()
	var topics []*pb.Topic
	for i := 1; i < 3; i++ {
		topics = append(topics, mustCreateTopic(t, pclient, &pb.Topic{
			Name:   fmt.Sprintf("projects/P/topics/T%d", i),
			Labels: map[string]string{"num": fmt.Sprintf("%d", i)},
		}))
	}
	if got, want := len(server.gServer.topics), len(topics); got != want {
		t.Fatalf("got %d topics, want %d", got, want)
	}
	for _, top := range topics {
		got, err := pclient.GetTopic(ctx, &pb.GetTopicRequest{Topic: top.Name})
		if err != nil {
			t.Fatal(err)
		}
		if !testutil.Equal(got, top) {
			t.Errorf("\ngot %+v\nwant %+v", got, top)
		}
	}

	res, err := pclient.ListTopics(ctx, &pb.ListTopicsRequest{Project: "projects/P"})
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res.Topics, topics; !testutil.Equal(got, want) {
		t.Errorf("\ngot %+v\nwant %+v", got, want)
	}

	for _, top := range topics {
		if _, err := pclient.DeleteTopic(ctx, &pb.DeleteTopicRequest{Topic: top.Name}); err != nil {
			t.Fatal(err)
		}
	}
	if got, want := len(server.gServer.topics), 0; got != want {
		t.Fatalf("got %d topics, want %d", got, want)
	}
}

func TestSubscriptions(t *testing.T) {
	pclient, sclient, server := newFake(t)
	defer server.Close()

	ctx := context.Background()
	topic := mustCreateTopic(t, pclient, &pb.Topic{Name: "projects/P/topics/T"})
	var subs []*pb.Subscription
	for i := 0; i < 3; i++ {
		subs = append(subs, mustCreateSubscription(t, sclient, &pb.Subscription{
			Name:               fmt.Sprintf("projects/P/subscriptions/S%d", i),
			Topic:              topic.Name,
			AckDeadlineSeconds: int32(10 * (i + 1)),
		}))
	}

	if got, want := len(server.gServer.subs), len(subs); got != want {
		t.Fatalf("got %d subscriptions, want %d", got, want)
	}
	for _, s := range subs {
		got, err := sclient.GetSubscription(ctx, &pb.GetSubscriptionRequest{Subscription: s.Name})
		if err != nil {
			t.Fatal(err)
		}
		if !testutil.Equal(got, s) {
			t.Errorf("\ngot %+v\nwant %+v", got, s)
		}
	}

	res, err := sclient.ListSubscriptions(ctx, &pb.ListSubscriptionsRequest{Project: "projects/P"})
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res.Subscriptions, subs; !testutil.Equal(got, want) {
		t.Errorf("\ngot %+v\nwant %+v", got, want)
	}

	res2, err := pclient.ListTopicSubscriptions(ctx, &pb.ListTopicSubscriptionsRequest{Topic: topic.Name})
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(res2.Subscriptions), len(subs); got != want {
		t.Fatalf("got %d subs, want %d", got, want)
	}
	for i, got := range res2.Subscriptions {
		want := subs[i].Name
		if !testutil.Equal(got, want) {
			t.Errorf("\ngot %+v\nwant %+v", got, want)
		}
	}

	for _, s := range subs {
		if _, err := sclient.DeleteSubscription(ctx, &pb.DeleteSubscriptionRequest{Subscription: s.Name}); err != nil {
			t.Fatal(err)
		}
	}
	if got, want := len(server.gServer.subs), 0; got != want {
		t.Fatalf("got %d subscriptions, want %d", got, want)
	}
}

func TestSubscriptionErrors(t *testing.T) {
	_, sclient, srv := newFake(t)
	defer srv.Close()

	ctx := context.Background()

	// TODO(jba): Go1.9: use t.Helper()
	checkCode := func(msg string, err error, want codes.Code) {
		if status.Code(err) != want {
			t.Errorf("%s: got %v, want code %s", msg, err, want)
		}
	}

	_, err := sclient.GetSubscription(ctx, &pb.GetSubscriptionRequest{})
	checkCode("GetSubscription", err, codes.InvalidArgument)
	_, err = sclient.GetSubscription(ctx, &pb.GetSubscriptionRequest{Subscription: "s"})
	checkCode("GetSubscription", err, codes.NotFound)
	_, err = sclient.UpdateSubscription(ctx, &pb.UpdateSubscriptionRequest{})
	checkCode("UpdateSubscription", err, codes.InvalidArgument)
	_, err = sclient.UpdateSubscription(ctx, &pb.UpdateSubscriptionRequest{Subscription: &pb.Subscription{}})
	checkCode("UpdateSubscription", err, codes.InvalidArgument)
	_, err = sclient.UpdateSubscription(ctx, &pb.UpdateSubscriptionRequest{Subscription: &pb.Subscription{Name: "s"}})
	checkCode("UpdateSubscription", err, codes.NotFound)
	_, err = sclient.DeleteSubscription(ctx, &pb.DeleteSubscriptionRequest{})
	checkCode("DeleteSubscription", err, codes.InvalidArgument)
	_, err = sclient.DeleteSubscription(ctx, &pb.DeleteSubscriptionRequest{Subscription: "s"})
	checkCode("DeleteSubscription", err, codes.NotFound)
	_, err = sclient.Acknowledge(ctx, &pb.AcknowledgeRequest{})
	checkCode("Acknowledge", err, codes.InvalidArgument)
	_, err = sclient.Acknowledge(ctx, &pb.AcknowledgeRequest{Subscription: "s"})
	checkCode("Acknowledge", err, codes.NotFound)
	_, err = sclient.ModifyAckDeadline(ctx, &pb.ModifyAckDeadlineRequest{})
	checkCode("ModifyAckDeadline", err, codes.InvalidArgument)
	_, err = sclient.ModifyAckDeadline(ctx, &pb.ModifyAckDeadlineRequest{Subscription: "s"})
	checkCode("ModifyAckDeadline", err, codes.NotFound)
	_, err = sclient.Pull(ctx, &pb.PullRequest{})
	checkCode("Pull", err, codes.InvalidArgument)
	_, err = sclient.Pull(ctx, &pb.PullRequest{Subscription: "s"})
	checkCode("Pull", err, codes.NotFound)
	_, err = sclient.Seek(ctx, &pb.SeekRequest{})
	checkCode("Seek", err, codes.InvalidArgument)
	srt := &pb.SeekRequest_Time{Time: ptypes.TimestampNow()}
	_, err = sclient.Seek(ctx, &pb.SeekRequest{Target: srt})
	checkCode("Seek", err, codes.InvalidArgument)
	_, err = sclient.Seek(ctx, &pb.SeekRequest{Target: srt, Subscription: "s"})
	checkCode("Seek", err, codes.NotFound)
}

func TestPublish(t *testing.T) {
	s := NewServer()
	defer s.Close()

	var ids []string
	for i := 0; i < 3; i++ {
		ids = append(ids, s.Publish("projects/p/topics/t", []byte("hello"), nil))
	}
	s.Wait()
	ms := s.Messages()
	if got, want := len(ms), len(ids); got != want {
		t.Errorf("got %d messages, want %d", got, want)
	}
	for i, id := range ids {
		if got, want := ms[i].ID, id; got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	m := s.Message(ids[1])
	if m == nil {
		t.Error("got nil, want a message")
	}
}

// Note: this sets the fake's "now" time, so it is sensitive to concurrent changes to "now".
func publish(t *testing.T, pclient pb.PublisherClient, topic *pb.Topic, messages []*pb.PubsubMessage) map[string]*pb.PubsubMessage {
	pubTime := time.Now()
	now.Store(func() time.Time { return pubTime })
	defer func() { now.Store(time.Now) }()

	res, err := pclient.Publish(context.Background(), &pb.PublishRequest{
		Topic:    topic.Name,
		Messages: messages,
	})
	if err != nil {
		t.Fatal(err)
	}
	tsPubTime, err := ptypes.TimestampProto(pubTime)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]*pb.PubsubMessage{}
	for i, id := range res.MessageIds {
		want[id] = &pb.PubsubMessage{
			Data:        messages[i].Data,
			Attributes:  messages[i].Attributes,
			MessageId:   id,
			PublishTime: tsPubTime,
		}
	}
	return want
}

func TestPull(t *testing.T) {
	pclient, sclient, srv := newFake(t)
	defer srv.Close()

	top := mustCreateTopic(t, pclient, &pb.Topic{Name: "projects/P/topics/T"})
	sub := mustCreateSubscription(t, sclient, &pb.Subscription{
		Name:               "projects/P/subscriptions/S",
		Topic:              top.Name,
		AckDeadlineSeconds: 10,
	})

	want := publish(t, pclient, top, []*pb.PubsubMessage{
		{Data: []byte("d1")},
		{Data: []byte("d2")},
		{Data: []byte("d3")},
	})
	got := pubsubMessages(pullN(t, len(want), sclient, sub))
	if diff := testutil.Diff(got, want); diff != "" {
		t.Error(diff)
	}

	res, err := sclient.Pull(context.Background(), &pb.PullRequest{Subscription: sub.Name})
	if err != nil {
		t.Fatal(err)
	}
	if len(res.ReceivedMessages) != 0 {
		t.Errorf("got %d messages, want zero", len(res.ReceivedMessages))
	}
}

func TestStreamingPull(t *testing.T) {
	// A simple test of streaming pull.
	pclient, sclient, srv := newFake(t)
	defer srv.Close()

	top := mustCreateTopic(t, pclient, &pb.Topic{Name: "projects/P/topics/T"})
	sub := mustCreateSubscription(t, sclient, &pb.Subscription{
		Name:               "projects/P/subscriptions/S",
		Topic:              top.Name,
		AckDeadlineSeconds: 10,
	})

	want := publish(t, pclient, top, []*pb.PubsubMessage{
		{Data: []byte("d1")},
		{Data: []byte("d2")},
		{Data: []byte("d3")},
	})
	got := pubsubMessages(streamingPullN(t, len(want), sclient, sub))
	if diff := testutil.Diff(got, want); diff != "" {
		t.Error(diff)
	}
}

func TestStreamingPullAck(t *testing.T) {
	// Ack each message as it arrives. Make sure we don't see dups.
	minAckDeadlineSecs = 1
	pclient, sclient, srv := newFake(t)
	defer srv.Close()

	top := mustCreateTopic(t, pclient, &pb.Topic{Name: "projects/P/topics/T"})
	sub := mustCreateSubscription(t, sclient, &pb.Subscription{
		Name:               "projects/P/subscriptions/S",
		Topic:              top.Name,
		AckDeadlineSeconds: 1,
	})

	_ = publish(t, pclient, top, []*pb.PubsubMessage{
		{Data: []byte("d1")},
		{Data: []byte("d2")},
		{Data: []byte("d3")},
	})

	got := map[string]bool{}
	spc := mustStartStreamingPull(t, sclient, sub)
	time.AfterFunc(time.Duration(3*minAckDeadlineSecs)*time.Second, func() {
		if err := spc.CloseSend(); err != nil {
			t.Errorf("CloseSend: %v", err)
		}
	})

	for {
		res, err := spc.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		req := &pb.StreamingPullRequest{}
		for _, m := range res.ReceivedMessages {
			if got[m.Message.MessageId] {
				t.Fatal("duplicate message")
			}
			got[m.Message.MessageId] = true
			req.AckIds = append(req.AckIds, m.AckId)
		}
		if err := spc.Send(req); err != nil {
			t.Fatal(err)
		}
	}
}

func TestAcknowledge(t *testing.T) {
	ctx := context.Background()
	pclient, sclient, srv := newFake(t)
	defer srv.Close()

	top := mustCreateTopic(t, pclient, &pb.Topic{Name: "projects/P/topics/T"})
	sub := mustCreateSubscription(t, sclient, &pb.Subscription{
		Name:               "projects/P/subscriptions/S",
		Topic:              top.Name,
		AckDeadlineSeconds: 10,
	})

	publish(t, pclient, top, []*pb.PubsubMessage{
		{Data: []byte("d1")},
		{Data: []byte("d2")},
		{Data: []byte("d3")},
	})
	msgs := streamingPullN(t, 3, sclient, sub)
	var ackIDs []string
	for _, m := range msgs {
		ackIDs = append(ackIDs, m.AckId)
	}
	if _, err := sclient.Acknowledge(ctx, &pb.AcknowledgeRequest{
		Subscription: sub.Name,
		AckIds:       ackIDs,
	}); err != nil {
		t.Fatal(err)
	}
	smsgs := srv.Messages()
	if got, want := len(smsgs), 3; got != want {
		t.Fatalf("got %d messages, want %d", got, want)
	}
	for _, sm := range smsgs {
		if sm.Acks != 1 {
			t.Errorf("message %s: got %d acks, want 1", sm.ID, sm.Acks)
		}
	}
}

func TestModAck(t *testing.T) {
	ctx := context.Background()
	pclient, sclient, srv := newFake(t)
	defer srv.Close()

	top := mustCreateTopic(t, pclient, &pb.Topic{Name: "projects/P/topics/T"})
	sub := mustCreateSubscription(t, sclient, &pb.Subscription{
		Name:               "projects/P/subscriptions/S",
		Topic:              top.Name,
		AckDeadlineSeconds: 10,
	})

	publish(t, pclient, top, []*pb.PubsubMessage{
		{Data: []byte("d1")},
		{Data: []byte("d2")},
		{Data: []byte("d3")},
	})
	msgs := streamingPullN(t, 3, sclient, sub)
	var ackIDs []string
	for _, m := range msgs {
		ackIDs = append(ackIDs, m.AckId)
	}
	if _, err := sclient.ModifyAckDeadline(ctx, &pb.ModifyAckDeadlineRequest{
		Subscription:       sub.Name,
		AckIds:             ackIDs,
		AckDeadlineSeconds: 0,
	}); err != nil {
		t.Fatal(err)
	}
	// Having nacked all three messages, we should see them again.
	msgs = streamingPullN(t, 3, sclient, sub)
	if got, want := len(msgs), 3; got != want {
		t.Errorf("got %d messages, want %d", got, want)
	}
}

func TestAckDeadline(t *testing.T) {
	// Messages should be resent after they expire.
	pclient, sclient, srv := newFake(t)
	defer srv.Close()

	minAckDeadlineSecs = 2
	top := mustCreateTopic(t, pclient, &pb.Topic{Name: "projects/P/topics/T"})
	sub := mustCreateSubscription(t, sclient, &pb.Subscription{
		Name:               "projects/P/subscriptions/S",
		Topic:              top.Name,
		AckDeadlineSeconds: minAckDeadlineSecs,
	})

	_ = publish(t, pclient, top, []*pb.PubsubMessage{
		{Data: []byte("d1")},
		{Data: []byte("d2")},
		{Data: []byte("d3")},
	})

	got := map[string]int{}
	spc := mustStartStreamingPull(t, sclient, sub)
	// In 5 seconds the ack deadline will expire twice, so we should see each message
	// exactly three times.
	time.AfterFunc(5*time.Second, func() {
		if err := spc.CloseSend(); err != nil {
			t.Errorf("CloseSend: %v", err)
		}
	})
	for {
		res, err := spc.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		for _, m := range res.ReceivedMessages {
			got[m.Message.MessageId]++
		}
	}
	for id, n := range got {
		if n != 3 {
			t.Errorf("message %s: saw %d times, want 3", id, n)
		}
	}
}

func TestMultiSubs(t *testing.T) {
	// Each subscription gets every message.
	pclient, sclient, srv := newFake(t)
	defer srv.Close()

	top := mustCreateTopic(t, pclient, &pb.Topic{Name: "projects/P/topics/T"})
	sub1 := mustCreateSubscription(t, sclient, &pb.Subscription{
		Name:               "projects/P/subscriptions/S1",
		Topic:              top.Name,
		AckDeadlineSeconds: 10,
	})
	sub2 := mustCreateSubscription(t, sclient, &pb.Subscription{
		Name:               "projects/P/subscriptions/S2",
		Topic:              top.Name,
		AckDeadlineSeconds: 10,
	})

	want := publish(t, pclient, top, []*pb.PubsubMessage{
		{Data: []byte("d1")},
		{Data: []byte("d2")},
		{Data: []byte("d3")},
	})
	got1 := pubsubMessages(streamingPullN(t, len(want), sclient, sub1))
	got2 := pubsubMessages(streamingPullN(t, len(want), sclient, sub2))
	if diff := testutil.Diff(got1, want); diff != "" {
		t.Error(diff)
	}
	if diff := testutil.Diff(got2, want); diff != "" {
		t.Error(diff)
	}
}

func TestMultiStreams(t *testing.T) {
	// Messages are handed out to the streams of a subscription in round-robin order.
	pclient, sclient, srv := newFake(t)
	defer srv.Close()

	top := mustCreateTopic(t, pclient, &pb.Topic{Name: "projects/P/topics/T"})
	sub := mustCreateSubscription(t, sclient, &pb.Subscription{
		Name:               "projects/P/subscriptions/S",
		Topic:              top.Name,
		AckDeadlineSeconds: 10,
	})
	want := publish(t, pclient, top, []*pb.PubsubMessage{
		{Data: []byte("d1")},
		{Data: []byte("d2")},
		{Data: []byte("d3")},
		{Data: []byte("d4")},
	})
	streams := []pb.Subscriber_StreamingPullClient{
		mustStartStreamingPull(t, sclient, sub),
		mustStartStreamingPull(t, sclient, sub),
	}
	got := map[string]*pb.PubsubMessage{}
	for i := 0; i < 2; i++ {
		for _, st := range streams {
			res, err := st.Recv()
			if err != nil {
				t.Fatal(err)
			}
			m := res.ReceivedMessages[0]
			got[m.Message.MessageId] = m.Message
		}
	}
	if diff := testutil.Diff(got, want); diff != "" {
		t.Error(diff)
	}
}

func TestStreamingPullTimeout(t *testing.T) {
	pclient, sclient, srv := newFake(t)
	defer srv.Close()

	timeout := 200 * time.Millisecond
	srv.SetStreamTimeout(timeout)
	top := mustCreateTopic(t, pclient, &pb.Topic{Name: "projects/P/topics/T"})
	sub := mustCreateSubscription(t, sclient, &pb.Subscription{
		Name:               "projects/P/subscriptions/S",
		Topic:              top.Name,
		AckDeadlineSeconds: 10,
	})
	stream := mustStartStreamingPull(t, sclient, sub)
	time.Sleep(2 * timeout)
	_, err := stream.Recv()
	if err != io.EOF {
		t.Errorf("got %v, want io.EOF", err)
	}
}

func TestSeek(t *testing.T) {
	pclient, sclient, srv := newFake(t)
	defer srv.Close()

	top := mustCreateTopic(t, pclient, &pb.Topic{Name: "projects/P/topics/T"})
	sub := mustCreateSubscription(t, sclient, &pb.Subscription{
		Name:               "projects/P/subscriptions/S",
		Topic:              top.Name,
		AckDeadlineSeconds: 10,
	})
	ts := ptypes.TimestampNow()
	_, err := sclient.Seek(context.Background(), &pb.SeekRequest{
		Subscription: sub.Name,
		Target:       &pb.SeekRequest_Time{Time: ts},
	})
	if err != nil {
		t.Errorf("Seeking: %v", err)
	}
}

func TestTryDeliverMessage(t *testing.T) {
	for _, test := range []struct {
		desc           string
		availStreamIdx int
		expectedOutIdx int
	}{
		{availStreamIdx: 0, expectedOutIdx: 0},
		// Stream 1 will always be marked for deletion.
		{availStreamIdx: 2, expectedOutIdx: 1}, // s0, s1 (deleted), s2, s3 becomes s0, s2, s3. So we expect outIdx=1.
		{availStreamIdx: 3, expectedOutIdx: 2}, // s0, s1 (deleted), s2, s3 becomes s0, s2, s3. So we expect outIdx=2.
	} {
		top := newTopic(&pb.Topic{Name: "some-topic"})
		sub := newSubscription(top, &sync.Mutex{}, &pb.Subscription{Name: "some-sub", Topic: "some-topic"})

		done := make(chan struct{}, 1)
		done <- struct{}{}
		sub.streams = []*stream{{}, {done: done}, {}, {}}

		msgc := make(chan *pb.ReceivedMessage, 1)
		sub.streams[test.availStreamIdx].msgc = msgc

		var d int
		idx, ok := sub.tryDeliverMessage(&message{deliveries: &d}, 0, time.Now())
		if !ok {
			t.Fatalf("[avail=%d]: expected msg to be put on stream %d's channel, but it was not", test.availStreamIdx, test.expectedOutIdx)
		}
		if idx != test.expectedOutIdx {
			t.Fatalf("[avail=%d]: expected msg to be put on stream %d, but it was put on %d", test.availStreamIdx, test.expectedOutIdx, idx)
		}
		select {
		case <-msgc:
		default:
			t.Fatalf("[avail=%d]: expected msg to be put on stream %d's channel, but it was not", test.availStreamIdx, idx)
		}
	}
}

func mustStartStreamingPull(t *testing.T, sc pb.SubscriberClient, sub *pb.Subscription) pb.Subscriber_StreamingPullClient {
	spc, err := sc.StreamingPull(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if err := spc.Send(&pb.StreamingPullRequest{Subscription: sub.Name}); err != nil {
		t.Fatal(err)
	}
	return spc
}

func pullN(t *testing.T, n int, sc pb.SubscriberClient, sub *pb.Subscription) map[string]*pb.ReceivedMessage {
	ctx := context.Background()
	got := map[string]*pb.ReceivedMessage{}
	for i := 0; len(got) < n; i++ {
		res, err := sc.Pull(ctx, &pb.PullRequest{Subscription: sub.Name, MaxMessages: int32(n - len(got))})
		if err != nil {
			t.Fatal(err)
		}
		for _, m := range res.ReceivedMessages {
			got[m.Message.MessageId] = m
		}
	}
	return got
}

func streamingPullN(t *testing.T, n int, sc pb.SubscriberClient, sub *pb.Subscription) map[string]*pb.ReceivedMessage {
	spc := mustStartStreamingPull(t, sc, sub)
	got := map[string]*pb.ReceivedMessage{}
	for i := 0; i < n; i++ {
		res, err := spc.Recv()
		if err != nil {
			t.Fatal(err)
		}
		for _, m := range res.ReceivedMessages {
			got[m.Message.MessageId] = m
		}
	}
	if err := spc.CloseSend(); err != nil {
		t.Fatal(err)
	}
	res, err := spc.Recv()
	if err != io.EOF {
		t.Fatalf("Recv returned <%v> instead of EOF; res = %v", err, res)
	}
	return got
}

func pubsubMessages(rms map[string]*pb.ReceivedMessage) map[string]*pb.PubsubMessage {
	ms := map[string]*pb.PubsubMessage{}
	for k, rm := range rms {
		ms[k] = rm.Message
	}
	return ms
}

func mustCreateTopic(t *testing.T, pc pb.PublisherClient, topic *pb.Topic) *pb.Topic {
	top, err := pc.CreateTopic(context.Background(), topic)
	if err != nil {
		t.Fatal(err)
	}
	return top
}

func mustCreateSubscription(t *testing.T, sc pb.SubscriberClient, sub *pb.Subscription) *pb.Subscription {
	sub, err := sc.CreateSubscription(context.Background(), sub)
	if err != nil {
		t.Fatal(err)
	}
	return sub
}

// Note: be sure to close server!
func newFake(t *testing.T) (pb.PublisherClient, pb.SubscriberClient, *Server) {
	srv := NewServer()
	conn, err := grpc.Dial(srv.Addr, grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	return pb.NewPublisherClient(conn), pb.NewSubscriberClient(conn), srv
}
