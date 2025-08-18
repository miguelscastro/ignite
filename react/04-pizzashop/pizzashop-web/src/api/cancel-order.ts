import { api } from '@/lib/axios'

export interface cancelOrderParams {
  orderId: string
}

export async function cancelOrder({ orderId }: cancelOrderParams) {
  const response = await api.patch(`/orders/${orderId}/cancel`)

  return response.data
}
