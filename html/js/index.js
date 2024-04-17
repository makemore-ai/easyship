// 获取searchArrow和searchInput元素
var searchArrow = document.getElementById('searchArrow');
var searchInput = document.getElementById('searchInput');
// 为searchArrow添加点击事件监听器
searchArrow.addEventListener('click', function () {
    this.classList.add('active'); // 点击时添加active类
    performSearch(); // 执行搜索函数
});

// 为searchInput添加回车事件监听器
searchInput.addEventListener('keydown', function (event) {
    if (event.key === 'Enter') { // 检查按键码是否为回车（键码13）
        event.preventDefault(); // 阻止默认的回车行为（如提交表单）
        searchArrow.classList.add('active'); // 添加active类以改变箭头颜色
        performSearch(); // 执行搜索函数
    }
});

var textCopyButtonList = document.getElementsByClassName('copy-button');
if (textCopyButtonList.length > 0) {
    for (var i = 0; i < textCopyButtonList.length; ++i) {
        let item = textCopyButtonList[i];
        console.log(item.id)
        item.addEventListener('click', function () {
            let buttonId = item.id
            let buttonInfos = buttonId.split('-');
            let textId = "text-" + buttonInfos[1]
            let textElement = document.getElementById(textId)

            navigator.clipboard.writeText(textElement.textContent.trim())
                .then(function () {
                    console.log('复制成功！');
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
// 执行搜索的函数
function performSearch() {
    var searchTerm = searchInput.value; // 获取输入框中的值
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
            console.log(!data['status'])
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
                    " <button id={{printf \"textCopyButton-${id}\"  .Id}} class=\"copy-button\">Copy</button>\n" +
                    "            <h3>${label_name}</h3>\n" +
                    "            <p id={{printf \"text-${id}\"  .Id}}>\n" +
                    "                ${prompt_zh}\n" +
                    "            </p>\n" +
                    "</div>"
                appendHtml = appendHtml.replace(/\$\{id}/, searchPromptData['id'])
                appendHtml = appendHtml.replace(/\$\{label_name}/, searchPromptData['label_name'])
                appendHtml = appendHtml.replace(/\$\{prompt_zh}/, searchPromptData['prompt_zh'])
                innerHtml += appendHtml
            }
            searchDiv.innerHTML = innerHtml
            console.log('Request list:', data['data'])
            console.log('Request successful:', data);
            // 在这里处理返回的数据
        })
        .catch(error => {
            console.error('网络错误：', error);
            // 在这里处理错误
        });
    searchArrow.classList.remove('active')
}
