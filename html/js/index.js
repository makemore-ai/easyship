// 获取searchArrow和searchInput元素
var searchArrow = document.getElementById('searchArrow');
var searchInput = document.getElementById('searchInput');
//默认不展示
document.getElementById('refresh-div').style.display = 'none';
// 为searchArrow添加点击事件监听器
searchArrow.addEventListener('click', function () {
    this.classList.add('active'); // 点击时添加active类
    doSearch()
});

document.getElementById('text-refresh').addEventListener('mouseleave', function (event) {
    console.log('mouse over')
    // 阻止默认滚动行为
    document.body.style.overflow = 'auto'
})

document.getElementById('text-refresh').addEventListener('mouseenter', function (event) {
    // 阻止默认滚动行为
    document.body.style.overflow = 'hidden'
})

// 为searchInput添加回车事件监听器
searchInput.addEventListener('keydown', function (event) {
    if (event.key === 'Enter') { // 检查按键码是否为回车（键码13）
        event.preventDefault(); // 阻止默认的回车行为（如提交表单）
        searchArrow.classList.add('active'); // 添加active类以改变箭头颜色
        doSearch()
    }
});

function doSearch() {
    let searchTerm = searchInput.value; // 获取输入框中的值
    if (searchTerm === undefined || searchTerm === '' || searchTerm == null) {
        // 展示推荐模块
        document.getElementById('refresh-div').style.display = '';
    } else {
        performRefresh(searchTerm)
        performSearch(searchTerm); // 执行搜索函数
        //隐藏推荐模块
        document.getElementById('recommend-div').style.display = 'None';
    }
}
function addButtonCopy() {
    var textCopyButtonList = document.getElementsByClassName('copy-button');
    if (textCopyButtonList.length > 0) {
        for (var i = 0; i < textCopyButtonList.length; ++i) {
            let item = textCopyButtonList[i];
            item.addEventListener('click', function () {
                let buttonId = item.id
                let buttonInfos = buttonId.split('-');
                let textId = "text-" + buttonInfos[1]
                let textElement = document.getElementById(textId)

                navigator.clipboard.writeText(textElement.textContent.trim())
                    .then(function () {
                        // 可以给用户一些反馈，比如改变按钮的文本或显示一个提示
                        item.textContent = 'Success';
                        // 可选：稍后将按钮文本恢复为原始值
                        setTimeout(function () {
                            item.textContent = 'Copy';
                        }, 1000); // 例如，1秒后恢复
                    })
                    .catch(function (err) {
                        // 复制失败，可能是因为用户拒绝了权限或者浏览器不支持该功能
                        console.error('复制失败:', err);
                        alert('无法复制文本，请检查你的浏览器设置或尝试手动复制。');
                    });
            });
        }
    }
}
addButtonCopy()
// 执行搜索的函数
function performSearch(searchTerm) {
    let searchDiv = document.getElementById('search-div');
    searchDiv.innerHTML = ""
    if(searchTerm.trim() === "") {
        return
    }
// 创建一个 FormData 实例
    var formData = new FormData();
    formData.set('searchText', searchTerm)

// 将字符串作为表单字段添加到 FormData 对象中
    fetch('/searchPrompt', {
        method: 'POST',
        body: formData
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Network Error');
            }
            return response.json()
        })
        .then(data => {
            if(data['status'] !== 0 || !data['data']) {
                throw new Error('System Error');
            }
            let resData = data['data']
            let innerHtml = "<p class = \"title-text\">搜索结果：</p>" + "\n"
            if (resData['search_prompt_list'].length === 0) {
                searchDiv.innerHTML = innerHtml
                return
            }
            // 最多展示8个
            for (var i = 0; i < resData['search_prompt_list'].length && i < 8; ++i) {
                let searchPromptData = resData['search_prompt_list'][i]
                let appendHtml = " <div class=\"content-module\">" +
                    " <button id=\"textCopyButton-${id}\" class=\"copy-button\">Copy</button>\n" +
                    "            <h3>${label_name}</h3>\n" +
                    "            <p id=\"text-${id}\">\n" +
                    "                ${prompt_zh}\n" +
                    "            </p>\n" +
                    "</div>"
                appendHtml = appendHtml.replaceAll(/\$\{id}/g, searchPromptData['id'])
                appendHtml = appendHtml.replace(/\$\{label_name}/, searchPromptData['label_name'])
                appendHtml = appendHtml.replace(/\$\{prompt_zh}/, searchPromptData['prompt_zh'])
                innerHtml += appendHtml
            }
            searchDiv.innerHTML = innerHtml
            // 在这里处理返回的数据
            addButtonCopy()
        })
        .catch(error => {
            console.error('网络错误：', error);
            // 在这里处理错误
        })

}

function performRefresh(searchTerm) {
    document.getElementById('refresh-div').value = '';
    document.getElementById('refresh-div').style.display = '';
    var refreshSSE = new EventSource('/refreshPrompt?prompt=' + searchTerm)
    refreshSSE.addEventListener("continue", function(e) {
        if (typeof e.data !== undefined && e.data !== "" &&  e.data != null) {
            handleSSEStreamData(refreshSSE, e.data)
        }
    })
    refreshSSE.addEventListener("stop", function(e) {
        if (typeof e.data !== undefined && e.data !== "" &&  e.data != null) {
            handleSSEStreamData(refreshSSE, e.data)
        }
        refreshSSE.close()
    })
    refreshSSE.onerror = function (e) {
        refreshSSE.close()
    }
}
function handleSSEStreamData(refreshSSE, data) {
    let strList =  splitByNewline(data)
    let dataStr = strList[1]
    let dataObj = JSON.parse(dataStr)
    // 错误 返回停止
    if(dataObj.status !== 0) {
        refreshSSE.close()
        return
    }
    let searchInput = document.getElementById('text-refresh');
    if (typeof dataObj.data === undefined || dataObj.data === "" || dataObj.data == null) {
        //空值不做处理
        return
    }
    searchInput.value = dataObj.data
}
function splitByNewline(str) {
    // 使用正则表达式匹配所有类型的换行符
    const regex = /\r\n|\n|\r/;
    return str.split(regex);
}
