import { http, HttpResponse } from 'msw'

import { type getMonthRevenueResponse } from '../get-month-revenue'

export const getMonthRevenueMock = http.get<
  never,
  never,
  getMonthRevenueResponse
>('/metrics/month-receipt', () => {
  return HttpResponse.json({
    receipt: 20000,
    diffFromLastMonth: 10,
  })
})
