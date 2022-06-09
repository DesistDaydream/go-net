// 绑定页面上的按钮元素
let StockIn = document.querySelector('button[name=StockIn]');
let StockOut = document.querySelector('button[name=StockOut]');
let StockQuery = document.querySelector('button[name=StockQuery]');

StockIn.onclick = function () {
    // 跳转到入库页面
    window.location.href = './stock-in.html';
}
StockOut.onclick = function () {
    // 跳转到出库页面
    window.location.href = './stock-out.html';
}
StockQuery.onclick = function () {
    // 跳转到查询页面
    window.location.href = './stock-query.html';
}