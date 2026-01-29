import { http, HttpResponse } from 'msw'

import { type GetManagedRestaurantResponse } from '../get-managed-restaurant'

export const getManagedRestaurantMock = http.get<
  never,
  never,
  GetManagedRestaurantResponse
>('/managed-restaurant', () => {
  return HttpResponse.json({
    id: 'custom-restaurant-id',
    name: 'Pizza Shop',
    managerId: 'custom-user-id',
    description: 'A place with delicious pizza',
    createdAt: new Date(),
    updatedAt: null,
  })
})
