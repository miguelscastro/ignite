import { http, HttpResponse } from 'msw'

import type { getDailyRevenueInPeriodResponse } from '../get-daily-revenue-in-period'

export const getDailyRevenueInPeriodMock = http.get<
  never,
  never,
  getDailyRevenueInPeriodResponse
>('/metrics/daily-receipt-in-period', () => {
  return HttpResponse.json([
    { date: '13/05/2024', receipt: 2350 },
    { date: '14/05/2024', receipt: 1890 },
    { date: '15/05/2024', receipt: 3120 },
    { date: '16/05/2024', receipt: 2500 },
    { date: '17/05/2024', receipt: 2880 },
    { date: '18/05/2024', receipt: 1950 },
    { date: '19/05/2024', receipt: 3010 },
  ])
})
