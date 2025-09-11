import { http, HttpResponse } from 'msw'

import type { getPopularProductsResponse } from '../get-popular-products'

export const getPopularProductsMock = http.get<
  never,
  never,
  getPopularProductsResponse
>('/metrics/popular-products', () => {
  return HttpResponse.json([
    { product: 'Calabresa', amount: 320 },
    { product: 'Mussarela', amount: 310 },
    { product: 'Frango com Catupiry', amount: 280 },
    { product: 'Marguerita', amount: 220 },
    { product: 'Quatro Queijos', amount: 200 },
  ])
})
