// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.635
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Base() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0\"><meta http-equiv=\"X-UA-Compatible\" content=\"ie=edge\"><title>KingsComp</title><script src=\"https://telegram.org/js/telegram-web-app.js\"></script><script defer src=\"https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js\"></script><link href=\"https://css.gg/css\" rel=\"stylesheet\"><link rel=\"stylesheet\" href=\"/static/fontiran.css?v=0.0.0-debug\"><style>\n        * {\n            padding: 0;\n            margin: 0;\n            box-sizing: border-box;\n            direction: rtl;\n            -webkit-user-select: none;\n            -moz-user-select: none;\n            -ms-user-select: none;\n            user-select: none;\n        }\n\n        i {\n            float: right;\n            margin-left: 4px;\n        }\n\n        .gg-loadbar {\n            padding-top: 8px;\n        }\n\n        body {\n            --bg-color: var(--tg-theme-bg-color);\n            background-color: var(--bg-color);\n            font-family: Ravi FaNum, serif;\n            color: var(--tg-theme-text-color);\n            margin: 48px 24px;\n            padding: 0;\n            color-scheme: var(--tg-color-scheme);\n        }\n\n        a, button, p, div, h1, h2, h3, h4, h5, h6, b, strong, em {\n            font-family: Ravi FaNum, serif;\n        }\n\n        a {\n            color: var(--tg-theme-link-color);\n        }\n\n        .spinner {\n            width: 56px;\n            height: 56px;\n            border-radius: 50%;\n            background: radial-gradient(farthest-side, var(--tg-theme-button-color) 94%, #0000) top/9px 9px no-repeat,\n            conic-gradient(#0000 30%, var(--tg-theme-button-color));\n            -webkit-mask: radial-gradient(farthest-side, #0000 calc(100% - 9px), #000 0);\n            animation: spinner-aib1d7 1s infinite linear;\n        }\n\n        .hint {\n            color: var(--tg-theme-hint-color);\n        }\n\n        @keyframes spinner-aib1d7 {\n            to {\n                transform: rotate(360deg);\n            }\n        }\n\n        .center {\n            display: flex;\n            align-items: center;\n            justify-content: center;\n            flex-direction: column;\n        }\n\n        .tg-button {\n            padding: 12px;\n            background: var(--tg-theme-button-color);\n            color: var(--tg-theme-button-text-color);\n            border: none;\n            outline: none;\n            cursor: pointer;\n            text-align: center;\n            text-decoration: none;\n            display: inline-block;\n            font-size: 16px;\n            border-radius: 8px;\n        }\n\n        [x-cloak] {\n            display: none !important;\n        }\n\n        .time-indicator {\n            -webkit-mask: linear-gradient(270deg, rgba(0, 0, 0, 1), transparent 80%);\n            height: 10px;\n            width: 80px;\n            background: var(--tg-theme-button-color);\n            color: var(--tg-theme-button-text-color);\n            -webkit-text-stroke: 2px var(--tg-theme-button-color);\n            transition: 0.2s;\n        }\n\n        .anim-fade-in {\n            animation: fadeInScale 0.4s cubic-bezier(0, 0.55, 0.45, 1) 0s 1 normal forwards;\n        }\n\n        @keyframes fadeInScale {\n            0% {\n                opacity: 0;\n                transform: scale(1.4);\n            }\n\n            100% {\n                opacity: 1;\n                transform: scale(1);\n            }\n        }\n\n        .time-indicator-holder {\n            width: 100%;\n            border: 2px var(--tg-theme-button-color) solid;\n            height: 14px;\n            direction: ltr !important;\n            overflow: hidden;\n            position: relative;\n            border-radius: 4px;\n        }\n\n        .time-indicator-text {\n            position: absolute;\n            left: 50%;\n            top: 50%;\n            transform: translateX(-50%) translateY(-50%);\n            color: var(--tg-theme-hint-color);\n            font-size: 16px;\n            font-weight: bold;\n        }\n\n        .box-with-border {\n            border: 1px var(--tg-theme-hint-color) solid;\n            border-radius: 4px;\n            width: 100%;\n            color: var(--tg-theme-text-color);\n            padding: 5px 10px;\n        }\n\n        .flex-row {\n            display: flex;\n            justify-content: space-between;\n            align-items: center;\n        }\n\n        .tg-button-bordered {\n            color: var(--tg-theme-text-color);\n            border: 2px var(--tg-theme-text-color) solid;\n            outline: none;\n            background: transparent;\n            padding: 5px 10px;\n            border-radius: 4px;\n        }\n    </style></head><body x-data=\"{\n    theme: window.Telegram.WebApp.themeParams,\n    auth: 0, // 0 means validating, 1 means validate, 2 means invalid,\n    authError: &#39;&#39;\n}\" x-init=\"\n        let response = await post(&#34;/auth/validate&#34;)\n\n        if (!response[&#34;data&#34;][&#34;is_valid&#34;]) {\n            auth = 2;\n            return\n        }\n\n        auth = 1;\n        WebApp.ready();\n        WebApp.expand();\n        WebApp.enableClosingConfirmation();\n\"><template x-if=\"auth === 0\"><div class=\"center\"><div><div class=\"spinner\"></div></div><p class=\"hint\">لطفا کمی صبر کنید</p></div></template><template x-if=\"auth === 1\" x-transition x-cloak>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</template><template x-if=\"auth === 2\" x-cloak><div class=\"center\"><h2>مشکلی پیش آمده :(</h2><p>لطفا دوباره تلاش کنید یا به پشتیبانی اطلاع دهید</p></div></template><script>\n    const WebApp = window.Telegram.WebApp;\n\n    async function post(path, data = {}) {\n        let response = await fetch(path, {\n            method: \"POST\",\n            headers: {\n                \"Content-Type\": \"application/json\",\n                \"Authorization\": WebApp.initData,\n            },\n            body: JSON.stringify(data)\n        })\n\n        response = await response.json()\n        return response\n    }\n</script><script>\n    document.addEventListener('alpine:init', () => {\n        Alpine.store('lobby', {\n            isInit: false,\n            lobbyId: \"\",\n            currentLobby: {},\n            currentPlayer: {},\n            lobbyHash: \"\",\n            questionTimer: 0,\n            timerPercent: 0,\n            questionDuration: 0,\n\n            async initLobby() {\n                if (this.isInit) {\n                    return\n                }\n\n                let response = await post('/lobby/' + this.lobbyId + '/ready')\n                if (response['code'] !== 200) {\n                    alert(response['data']);\n                    WebApp.close()\n                }\n\n                this.currentPlayer = response[\"data\"];\n\n                response = await post('/lobby/' + this.lobbyId + '/info')\n                if (response['code'] !== 200) {\n                    alert(response['data']);\n                    WebApp.close()\n                }\n\n                this.currentLobby = response['data'];\n                this.runEventWorker()\n                this.questionTimerWorker()\n                this.isInit = true\n            },\n            setLobbyId(lobbyId) {\n                this.lobbyId = lobbyId\n            },\n            async runEventWorker() {\n                while (true) {\n                    let response = await this.readEvents()\n                    if (!response[\"ok\"]) {\n                        alert(response[\"data\"])\n                        return\n                    }\n                    this.currentLobby = response[\"data\"][\"lobby\"];\n                    this.lobbyHash = response[\"data\"][\"hash\"];\n                    if (this.currentLobby['state'] === 'ended') {\n                        return\n                    }\n                }\n            },\n            async readEvents() {\n                return await post('/lobby/' + this.lobbyId + '/events', {\n                    hash: this.lobbyHash,\n                })\n            },\n            async answered(answer) {\n                let currentQuestion = this.currentLobby.gameInfo.currentQuestion.index\n                await post(\"/lobby/\" + this.lobbyId + \"/answer\", {\n                    index: currentQuestion,\n                    answer: answer,\n                })\n            },\n            questionTimerWorker() {\n                setInterval(() => {\n                    let duration = this.currentLobby.gameInfo.questionEndsAt - this.currentLobby.gameInfo.questionStartedAt\n                    if (duration === 0) {\n                        return\n                    }\n                    let currentUnix = (new Date() / 1000)\n                    let newPercent = Math.floor((currentUnix - this.currentLobby.gameInfo.questionStartedAt) / duration * 100)\n                    if (newPercent >= 100)\n                        newPercent = 100\n                    this.timerPercent = newPercent\n                }, 300)\n            }\n        })\n    })\n</script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
