package multiagent

import (
	"fmt"

	"github.com/cloudwego/eino/schema"
)

// toolExecutionRetryHint returns a user message appended to the conversation to prompt
// the LLM to adjust after a tool execution error (tool not found, binary missing,
// runtime failure, network error, etc.).
func toolExecutionRetryHint() *schema.Message {
	return schema.UserMessage(`[System] Your previous tool call failed. Possible causes:
- The tool or sub-agent name does not exist (typo or unregistered name).
- The tool call arguments were not valid JSON.
- The tool's underlying binary is not installed or not in PATH.
- The tool encountered a runtime error (timeout, network failure, permission denied, etc.).

Please review the error message above, check available tools, and either:
1. Retry with corrected arguments or a different tool, OR
2. Inform the user about the limitation and proceed with an alternative approach.

[系统提示] 上一次工具调用失败，可能原因：
- 工具名或子代理名称不存在（拼写错误或未注册）；
- 工具调用参数不是合法 JSON；
- 工具依赖的底层二进制程序未安装或不在 PATH 中；
- 工具运行时遇到错误（超时、网络故障、权限不足等）。

请根据上述错误信息检查可用工具，然后：
1. 修正参数或改用其他工具重试，或者
2. 告知用户当前限制并采用替代方案继续。`)
}

// toolExecutionRecoveryTimelineMessage returns a message for the eino_recovery event
// displayed in the UI timeline when a tool execution error triggers a retry.
func toolExecutionRecoveryTimelineMessage(attempt int) string {
	return fmt.Sprintf(
		"工具调用执行失败。已向对话追加纠错提示并要求模型调整策略。"+
			"当前为第 %d/%d 轮完整运行。\n\n"+
			"Tool call execution failed. "+
			"A corrective hint was appended. This is full run %d of %d.",
		attempt+1, maxToolCallRecoveryAttempts, attempt+1, maxToolCallRecoveryAttempts,
	)
}
