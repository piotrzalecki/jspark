package main

import (
	"context"
	"encoding/json"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func callGPT(input Input) (completionObject, error) {
	var co completionObject
	llm, err := openai.New()
	if err != nil {
		return co, err
	}
	ctx := context.Background()
	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, 
			`You are engineer assistants who creates jira ticket content for a given prompt. \\
			Based on prompt between '<<' and '>>' create 'title' and 'summary' for jira ticket. \\
			Title have to be no longer than 7 words. \\
			Summary have to be short and consistent. \\
			Summary describes problem, don't propose solution. \\
			From now on respond in json format {"title": "generated title", "summary": "generated summary", "comments": "other response text you want to pass"}. This is mandatory`),
		llms.TextParts(llms.ChatMessageTypeHuman, "user <<Seems like our cooked db's are executing queries very slow.>>"),
		llms.TextParts(llms.ChatMessageTypeHuman, ` assistant 
		{
			"title": "Cooked DB Slow Queries",
			"summary": "Our cooked databases are experiencing slow query execution times. This ticket is to investigate issue. For a found issue remediation should be implemented and/or follow up tickets created.",
			"comments": "ticket created"
			}`),
		llms.TextParts(llms.ChatMessageTypeHuman, ``),
		llms.TextParts(llms.ChatMessageTypeHuman, "user <<Hi Piotr, the pods are able to connecto to neo4j (great). However, I notice it didn’t relaunch the pods in the ray cluster when I ran the prod pipeline. I had to change it again so that the current deployment ran fine, and I just manually deleted the head and worker pods (just then the ray cluster pulled the new image).>>"),
		llms.TextParts(llms.ChatMessageTypeHuman, `assistant
		{
			"title": "Pipeline doesn't restart Ray pods",
			"summary": "Investigate Ray pipeline to ensure pods are restarted after deployment to pick up latest image.",
			"comments": "ticket created"
			}`),
		llms.TextParts(llms.ChatMessageTypeHuman, "user <<fix terraform https://github.com/cloudreach/cmz-etl-debezium/blob/main/terraform/infrastructure/modules/cmz-etl-eks/main.tf#L150; set image to be persistent; using 'data' is causing node template updates when new image released what updates nodes and pods need to be moved what can cause down time>>"),
		llms.TextParts(llms.ChatMessageTypeHuman, `assistant 
		{
			"title": "Set persistent image in node template.",
			"summary": "This is to change terraform code in https://github.com/cloudreach/cmz-etl-debezium/blob/main/terraform/infrastructure/modules/cmz-etl-eks/main.tf#L150. Currently we are using 'data' and due to this node template is updated every time image is changing. This causes creation of new node pools and migrating pods from old to new node poll what can cause downtime.",
			"comments": "ticket created"
			}`),
			llms.TextParts(llms.ChatMessageTypeHuman, "user <<execute EKS upgrade from 1.26 to 1.29; process documented on https://cloudreach-software.atlassian.net/wiki/spaces/SRE/pages/37342674947/EKS+1.26+to+1.29+Upgrade+Plan; research done on CIS-3452>>"),
			llms.TextParts(llms.ChatMessageTypeHuman, `assistant 
			{
				"title": "EKS Upgrade 1.26 to 1.29",
				"summary": "This is to execute EKS cluster upgrade based on a plan documented on https://cloudreach-software.atlassian.net/wiki/spaces/SRE/pages/37342674947/EKS+1.26+to+1.29+Upgrade+Plan. Pre upgrade research has been already done and documented in a ticket CIS-3452",
				"comments": "ticket created"
				}`),
			llms.TextParts(llms.ChatMessageTypeHuman, "user <<" + input.Prompt + ">>"),
	}

	completion, err := llm.GenerateContent(ctx, content, llms.WithTemperature(0.1))

	if err != nil {
		return co, err
	}

	err = json.Unmarshal([]byte(completion.Choices[0].Content), &co)
	if err != nil {
		return co, err
	}
	return co, nil
}
