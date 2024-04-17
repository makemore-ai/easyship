package config

import (
	"context"
	"github.com/easyship/model/vo"
	"github.com/easyship/util"
)

var RECOMMEND_PROMPT_LSIT = []*vo.PromptVO{
	{
		Id:        1,
		LabelName: "写作助理",
		PromptZh: util.StrPtr("作为写作改进助理，您的任务是改进所提供文本的拼写、语法、清晰度、简洁性和整体可读性，同时分解长句子，" +
			"减少重复，并提供改进建议。请仅提供更正后的中文文本，并避免包含解释。请首先编辑以下文本：[文章内容]"),
		PromptEn: util.StrPtr("As a writing improvement assistant, your task is to improve the spelling, grammar, clarity, concision," +
			" and overall readability of the text provided, while breaking down long sentences, reducing repetition, and providing suggestions for improvement. " +
			"Please provide only the corrected Chinese version of the text and avoid including explanations. Please begin by editing the following text: [文章内容]"),
	},
	{
		Id:        2,
		LabelName: "周报生成器",
		PromptZh: util.StrPtr("使用下面提供的文本作为每周报告的基础，生成突出显示最重要点的简明摘要。报告应以 Markdown 格式编写，" +
			"并且应易于普通受众阅读和理解。特别是，专注于提供对利益相关者和决策者有用的见解和分析。您" +
			"还可以根据需要使用任何其他信息或来源。整个对话和说明均应以中文提供。请首先编辑以下文本：[工作内容]"),
		PromptEn: util.StrPtr("\"Using the provided text below as the basis for a weekly report, generate a concise " +
			"summary that highlights the most important points. The report should be written in markdown format and should " +
			"be easily readable and understandable for a general audience. In particular, focus on providing insights and " +
			"analysis that would be useful to stakeholders and decision-makers. You may also use any additional information " +
			"or sources as necessary. The entire conversation and instructions should be provided in Chinese. " +
			"Please begin by editing the following text: [工作内容]"),
	},
	{
		Id:        3,
		LabelName: "编程问题",
		PromptZh: util.StrPtr("我希望你担任 stackoverflow 的帖子回复工作。我会问与编程相关的问题，你会回答应该是什么答案。" +
			"我希望你只回复给定的答案，并在没有足够细节时写下解释。不要写解释。当我需要用告诉你一些事情时，我会将文本放在大括号内{像这样}。" +
			"整个对话和说明均应以中文提供。我的第一个问题是[编程问题]"),
		PromptEn: util.StrPtr("I want you to act as a stackoverflow post. I will ask programming-related questions and " +
			"you will reply with what the answer should be. I want you to only reply with the given answer, and write explanations " +
			"when there is not enough detail. do not write explanations. When I need to tell you something in English," +
			" I will do so by putting text inside curly brackets {like this}. The entire conversation and instructions should be provided in Chinese." +
			" My first question is [编程问题]"),
	},
	{
		Id:        4,
		LabelName: "论文提示",
		PromptZh: util.StrPtr("我想让你担任院士。您将负责研究您选择的主题并以论文或文章的形式展示研究结果。您的任务是确定可靠的来源，" +
			"以结构良好的方式组织材料，并通过引用准确记录材料。整个对话和说明均应以中文提供。我的第一个建议请求是[论文主题]"),
		PromptEn: util.StrPtr("I want you to act as an academician. You will be responsible for researching a topic " +
			"of your choice and presenting the findings in a paper or article form. Your task is to identify reliable sources, " +
			"organize the material in a well-structured way and document it accurately with citations. The entire conversation " +
			"and instructions should be provided in Chinese. My first suggestion request is [论文主题]"),
	},
}

// GetRecommendPromptList 获取推荐列表
func GetRecommendPromptList(ctx context.Context) []*vo.PromptVO {
	return RECOMMEND_PROMPT_LSIT
}
