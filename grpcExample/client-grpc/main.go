/*
 * Copyright 2019 American Express Travel Related Services Company, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 */

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/roost-io/roost-example/grpcExample/api"

	zb "github.com/ZB-io/zbio/client"
	zbutil "github.com/roost-io/roost-example/grpcExample/message"

	"google.golang.org/grpc"
)

func main() {
	// Command-Line Flag (Defines a string flag with specified name, default value, and usage string).
	// Default ip and port: 127.0.0.1:3000
	// Usage: ./client -ip=127.0.0.1:3000
	ipPtr := flag.String("ip", "127.0.0.1:3000", "Description: ip address")

	// Once all flags are declared, we call `flag.Parse()' to execute the command-line parsing.
	flag.Parse()

	// Connection is insecure because we are not using https.
	conn, err := grpc.Dial(*ipPtr, grpc.WithInsecure())
	if err != nil {
		// Fatalf kills the program after the error.
		log.Fatalf("❌ Error connecting to server %s. %v", *ipPtr, err)
	}
	log.Println("🚀 Launching client...")

	// defer is used to ensure that a function call is performed later in a program's execution, usually for purposes of cleanup.
	// Call to properly close the connection when the function returns.
	defer conn.Close()

	// Create a client for the the ProcessText service.
	client := api.NewProcessTextClient(conn)

	// Create a random client number between 1 to 10000.
	randomClientName := strconv.Itoa(seededRand.Intn(10000))

	log.Printf("💡 Client's name is %s", randomClientName)

	zbConfig := zbutil.Config(randomClientName)
	zbutil.InitZBIO(zbConfig)

	message := zb.Message{
		TopicName:     zbutil.TopicName, // default topicName
		Data:          []byte(fmt.Sprintf("grpc-client starting from grpcExample. zbClientName: %s\n", randomClientName)),
		HintPartition: "",
	}
	zbutil.SendMessageToZBIO([]zb.Message{message})

	for {

		var message zb.Message
		// Create a random string of length 10 to send to the server.
		randomMessage := randomString(10)

		// Create a context.
		ctx := context.Background()

		// Send a request.
		reqMessage := &api.InputRequest{
			Text:       randomMessage,
			ClientName: randomClientName}

		requestLog := fmt.Sprintf("⬅️ Client sent a message to server : %s", randomMessage)
		log.Println(requestLog)

		// Send requested message to zbio
		message = zb.Message{
			TopicName:     zbutil.TopicName, // default topicName
			Data:          []byte(requestLog),
			HintPartition: "",
		}
		zbutil.SendMessageToZBIO([]zb.Message{message})

		// Receive response.
		resp, err := client.Upper(ctx, reqMessage)
		if err != nil {
			log.Printf("❌ Error doing upper : %v", err)
		}

		responseLog := fmt.Sprintf("➡️ Received Response from server %v : %s ", resp.GetServerName(), resp.Text)
		log.Printf(responseLog)

		// Send response message to zbio
		message = zb.Message{
			TopicName:     zbutil.TopicName, // default topicName
			Data:          []byte(responseLog),
			HintPartition: "",
		}
		zbutil.SendMessageToZBIO([]zb.Message{message})

		// Sleep for 2 seconds before sending another message.
		time.Sleep(2 * time.Second)
	}
}
