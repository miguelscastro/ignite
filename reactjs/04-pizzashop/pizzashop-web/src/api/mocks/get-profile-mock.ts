import { http, HttpResponse } from 'msw'

import { type getProfileResponse } from '../get-profile'

export const getProfileMock = http.get<never, never, getProfileResponse>(
  '/me',
  () => {
    return HttpResponse.json({
      id: 'custom-user-id',
      name: 'John Doe',
      email: 'johndoe@example.com',
      phone: '312123223',
      role: 'manager',
      createdAt: new Date(),
      updatedAt: null,
    })
  },
)
