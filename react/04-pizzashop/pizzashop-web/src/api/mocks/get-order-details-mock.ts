import { http, HttpResponse } from 'msw'

import {
  type getOrderDetailsParams,
  type GetOrderDetailsResponse,
} from '../get-order-details'

export const getOrderDetailsMock = http.get<
  getOrderDetailsParams,
  never,
  GetOrderDetailsResponse
>('/orders/:orderId', ({ params }) => {
  return HttpResponse.json({
    id: params.orderId,
    customer: {
      name: 'John Doe',
      email: 'johndoe@example.com',
      phone: '12312321',
    },
    status: 'pending',
    createdAt: new Date().toISOString(),
    orderItems: [
      {
        id: 'order-item-1',
        priceInCents: 1000,
        product: { name: 'Pizza pepperoni' },
        quantity: 1,
      },
      {
        id: 'order-item-2',
        priceInCents: 3420,
        product: { name: 'Pizza marguerita' },
        quantity: 3,
      },
    ],
    totalInCents: 4420, // Exemplo: (1 * 1000) + (3 * 3420)
  })
})
