// 获取表单中的按钮
let stockQuery = document.querySelector('button[name=StockQuery]');
let stockReturn = document.querySelector('button[name=StockReturn]');

// 发送 ajax 请求,从 /api/query 中获取数据
let xhr = new XMLHttpRequest()
xhr.open("GET", "http://localhost:18080/api/query", true)
xhr.onload = function () {
    console.log(xhr.responseText)
    // 将响应数据填充到表格中
    let resp = JSON.parse(xhr.responseText)
    console.log(resp)
    let tbody = document.getElementById("refuge-tbody")
    for (let i = 0; i < resp.length; i++) {
        let tr = document.createElement("tr")
        let td1 = document.createElement("td")
        let td2 = document.createElement("td")
        let td3 = document.createElement("td")
        let td4 = document.createElement("td")
        let td5 = document.createElement("td")
        td1.innerText = resp[i].Provider
        td2.innerText = resp[i].Product
        td3.innerText = resp[i].Type
        td4.innerText = resp[i].Size
        td5.innerText = resp[i].Amount
        tr.appendChild(td1)
        tr.appendChild(td2)
        tr.appendChild(td3)
        tr.appendChild(td4)
        tr.appendChild(td5)
        tbody.appendChild(tr)
    }
}
xhr.send()

// 点击返回按钮，跳转到 order.html
stockReturn.onclick = function () {
    window.location.href = './order.html';
}