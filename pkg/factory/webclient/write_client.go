package webclient

import (
	"fmt"
	"strings"

	"github.com/jobergner/backent-cli/pkg/ast"
	. "github.com/jobergner/backent-cli/pkg/factory/utils"
	. "github.com/jobergner/backent-cli/pkg/typescript"
)

func (s *Factory) clientActions() []string {
	var methodString []string
	s.rangeActions(func(action ast.Action) *Code {
		var params []ast.Field
		var contentStrings []string
		rangeParams(action, func(param ast.Field) Param {
			params = append(params, param)
			contentStrings = append(contentStrings, param.Name)
			return Param{}
		})

		paramString := paramsStringTemplate(params)
		content := strings.Join(contentStrings, ", ")

		if action.Response == nil {
			s := methodTemplate(action.Name, paramString, content)
			methodString = append(methodString, s)
		} else {
			s := methodWithResponseTemplate(action.Name, paramString, content, Title(action.Name)+"Response")
			methodString = append(methodString, s)
		}

		return nil
	})

	return methodString
}

func methodTemplate(name, params, content string) string {
	return fmt.Sprintf(`  public %s(%s) {
    const messageID = generateID();
    const message: WebSocketMessage = {
      id: messageID,
      kind: MessageKind.Action%s,
      content: JSON.stringify({%s}),
    };
    setTimeout(() => {
      this.ws.send(JSON.stringify(message));
    }, 0);
  }`,
		name,
		params,
		Title(name),
		content,
	)
}

func paramsStringTemplate(params []ast.Field) string {
	var paramStrings []string
	for _, p := range params {
		paramStrings = append(paramStrings, fmt.Sprintf("%s: %s", p.Name, goTypeToTypescriptType(p.ValueString)))
	}

	return strings.Join(paramStrings, ", ")
}

func methodWithResponseTemplate(name, params, content, responseType string) string {
	return fmt.Sprintf(`  public %s(%s): Promise<%s> {
    const messageID = generateID();
    const message: WebSocketMessage = {
      id: messageID,
      kind: MessageKind.Action%s,
      content: JSON.stringify({%s}),
    };
    setTimeout(() => {
      this.ws.send(JSON.stringify(message));
    }, 0);
    return new Promise((resolve, reject) => {
      this.responseEmitter.on(messageID, (response: %s) => {
        resolve(response);
      });
      setTimeout(() => {
        reject(ErrResponseTimeout);
      }, responseTimeout);
    });
  }`,
		name,
		params,
		responseType,
		Title(name),
		content,
		responseType,
	)
}

func clientTemplate(methods []string) string {
	methodsString := strings.Join(methods, "\n")

	return fmt.Sprintf(`export class Client {
  private ws: WebSocket;
  private responseEmitter: EventEmitter;
  private id: string;
  constructor(url: string) {
    this.id = "";
    this.ws = new WebSocket(url);
    this.responseEmitter = new EventEmitter();
    this.ws.addEventListener("open", () => {
      console.log("WebSocket connection opened");
    });
    this.ws.addEventListener("message", (event) => {
      const message = JSON.parse(event.data);
      switch (message.kind) {
        case MessageKind.ID:
          this.id = message.content;
          break;
        case MessageKind.Update:
        case MessageKind.CurrentState:
          RECEIVEUPDATE(message.content);
          break;
        case MessageKind.Error:
          console.log(message);
          break;
        default:
          this.responseEmitter.emit(message.id, JSON.parse(message.content));
          break;
      }
    });
    this.ws.addEventListener("close", () => {
      console.log("WebSocket connection closed");
    });
  }
  public getID() {
    return this.id;
  }
%s
}`,
		methodsString,
	)
}

func (s *Factory) writeClient() *Factory {
	methodStrings := s.clientActions()
	client := clientTemplate(methodStrings)
	s.file.WriteString(client)
	return s
}
