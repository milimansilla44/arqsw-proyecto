export const CrearOrden = async () => {
  try {
    let cookie = Cookie.get("username");
    let id_user = parseInt(cookie.split(",")[0]);
    let orderDetail = { 'product_id': 0, 'quantity': 0, 'unit_price': 0 };
    let ordersDetail = [];

    let a = Cookie.get("cart" + id_user).split(";");

    for (let i = 0; i < a.length; i++) {
      let item = a[i];
      if (item != "") {
        let array = item.split(",");
        let product_id = parseInt(array[0]);
        let quantity = parseInt(array[1]);

        // Obtener información del producto, incluyendo el precio unitario
        let product = await GetProductById(product_id);
        let unit_price = product.product_unit_price;

        orderDetail = {
          'product_id': product_id,
          'quantity': quantity,
          'unit_price': unit_price
        };

        ordersDetail.push(orderDetail);
      }
    }

    const crearOrden = async () => {
      fetch('http://localhost:8090/neworder', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          user_id: id_user,
          details: ordersDetail
        })
      })
        .then(response => {
          if (response.status !== 201) {
            swal.fire({
              icon: 'error',
              text: "No se pudo realizar la compra"
            }).then((result) => {
              if (result.isConfirmed) {
                window.location.reload();
                return response.json();
              }
            });
          } else {
            swal.fire({
              icon: 'success',
              text: "Compra realizada con éxito"
            }).then((result) => {
              if (result.isConfirmed) {
                window.location.replace("/inicio");
                vaciarCarrito();
                return response.json();
              }
            });
          }
        });
    };

    crearOrden();
  } catch (error) {
    console.error("Error al crear la orden:", error);
  }
};
