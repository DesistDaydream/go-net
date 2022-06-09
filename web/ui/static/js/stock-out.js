// 获取页面上的元素
let FormEle = document.querySelector('form')
let CustomerEle = FormEle.querySelector('input[name=customer]')
let ProductEle = FormEle.querySelector('select[name=product]')
let SizeEle = FormEle.querySelector('select[name=size]')
let AmountEle = FormEle.querySelector('input[name=amount]')
let TypeEle = FormEle.querySelector('input[name=type]')
let InterestEle = FormEle.querySelectorAll('input[name=interest]')
// 获取出库按钮元素
let StockOutEle = FormEle.querySelector('button[name=StockOut]')
let BackEle = FormEle.querySelector('button[name=Back]')

// 点击提交按钮，提交表单
StockOutEle.onclick = function (e) {
    // 阻止默认行为
    e.preventDefault()

    // 获取表单数据
    let Customer = CustomerEle.value
    let Product = ProductEle.value
    let Size = SizeEle.value
    let Amount = AmountEle.value
    let Type = TypeEle.value
    // 转为 JSON
    let data = {
        customer: Customer,
        product: Product,
        type: Type,
        size: Size,
        amount: parseInt(Amount)
    }

    console.log(data)

    // 创建请求
    let xhr = new XMLHttpRequest()
    xhr.open('POST', 'http://localhost:18080/api/stock-out')
    xhr.onload = function () {
        // 响应数据 =ok 时出库成功
        if (xhr.responseText === 'ok') {
            alert('出库成功')
        } else {
            alert('出库失败')
        }
    }
    xhr.send(JSON.stringify(data))
}

// 点击返回时，返回上一页
BackEle.onclick = function (e) {
    // 阻止表单的默认提交行为
    e.preventDefault()
    window.location.href = "./order.html"
}