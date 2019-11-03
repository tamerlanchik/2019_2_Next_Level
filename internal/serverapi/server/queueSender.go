package server

import (
	"2019_2_Next_Level/internal/post"
	pb "2019_2_Next_Level/internal/post/messagequeue/service"
	"2019_2_Next_Level/internal/serverapi/server/config"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type QueueClient struct {
	queue      pb.MessageQueueClient
	Connection *grpc.ClientConn
}

func (q *QueueClient) Init() {
	var err error
	q.Connection, err = grpc.Dial(config.Conf.PostServiceHost+config.Conf.PostServiceSendPort, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Cannot connect to service")
		return
	}

	q.queue = pb.NewMessageQueueClient(q.Connection)
}

func (q *QueueClient) Destroy() {
	q.Connection.Close()
}

func (q *QueueClient) Put(email post.Email) error {
	p := (&ParcelAdapter{}).FromEmail(&email)
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	ctx := context.Background()
	_, err := q.queue.Enqueue(ctx, &p)
	return err
}

func (q *QueueClient) Get() (post.Email, error) {
	data, err := q.queue.Dequeue(context.Background(), &pb.Empty{})
	if err != nil {
		fmt.Println("Nil value")
		return post.Email{}, err
	}

	return (&ParcelAdapter{}).ToEmail(data), nil
}

type ParcelAdapter struct {
}

func (a *ParcelAdapter) ToEmail(from *pb.Email) post.Email {
	return post.Email{from.From, from.To, from.Body}
}

func (a *ParcelAdapter) FromEmail(from *post.Email) pb.Email {
	return pb.Email{
		From: from.From,
		To:   from.To,
		Body: from.Body,
	}
}