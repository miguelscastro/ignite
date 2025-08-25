import { api } from '@/lib/axios'

export interface getMonthCanceledOrdersAmountResponse {
  amount: number
  diffFromLastMonth: number
}

export async function getMonthCanceledOrdersAmount() {
  const response = await api.get<getMonthCanceledOrdersAmountResponse>(
    '/metrics/month-canceled-orders-amount',
  )

  return response.data
}
