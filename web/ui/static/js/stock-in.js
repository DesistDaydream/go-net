// 获取页面上的元素
let FormEle = document.querySelector('form')
let ProviderEle = FormEle.querySelector('input[name=provider]')
let ProductEle = FormEle.querySelector('select[name=product]')
let SizeEle = FormEle.querySelector('select[name=size]')
let AmountEle = FormEle.querySelector('input[name=amount]')
let TypeEle = FormEle.querySelector('input[name=type]')
// 获取入库按钮元素
let StockInEle = FormEle.querySelector('button[name=StockIn]')
let BackEle = FormEle.querySelector('button[name=Back]')

// 点击提交按钮，提交表单
StockInEle.onclick = function (e) {
    // 阻止表单的默认提交行为
    e.preventDefault()

    // 获取表单中的数据
    let Provider = ProviderEle.value
    let Product = ProductEle.value
    let Size = SizeEle.value
    let Amount = AmountEle.value
    let Type = TypeEle.value
    // 转为 JSON
    let data = {
        provider: Provider,
        product: Product,
        type: Type,
        size: Size,
        amount: parseInt(Amount)
    }

    console.log(data)

    // 验证stockInInputValue是否
    if (isNaN(Amount)) {
        alert('请输入自然数')
        return
    }

    let xhr = new XMLHttpRequest()
    xhr.open('POST', 'http://localhost:18080/api/stock-in', true)
    xhr.onload = function () {
        // 响应数据 =ok 时入库成功
        if (xhr.responseText === 'ok') {
            alert('入库成功')
        } else {
            alert('入库失败')
        }
    }
    // 发送请求，不改变 data 中字段的原始类型
    xhr.send(JSON.stringify(data))
}

// 点击返回时，返回上一页
BackEle.onclick = function (e) {
    // 阻止表单的默认提交行为
    e.preventDefault()
    window.location.href = "./order.html"
}