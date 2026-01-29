import { http, HttpResponse } from 'msw'

import { type getDayOrdersAmountResponse } from '../get-day-orders-amount'

export const getDayOrdersAmountMock = http.get<
  never,
  never,
  getDayOrdersAmountResponse
>('/metrics/day-orders-amount', () => {
  return HttpResponse.json({
    amount: 20,
    diffFromYesterday: -5,
  })
})
