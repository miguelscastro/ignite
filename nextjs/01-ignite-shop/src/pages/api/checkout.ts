import { stripe } from '@/lib/stripe'
import type { NextApiRequest, NextApiResponse } from 'next'

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse,
) {
  const { cartDetails } = req.body

  if (req.method !== 'POST') {
    return res.status(405).json({ error: 'Method not allowed' })
  }

  if (!cartDetails || Object.keys(cartDetails).length === 0) {
    return res.status(400).json({ error: 'Cart is empty or not found' })
  }

  try {
    const successUrl = `${process.env.NEXT_URL}/success?session_id={CHECKOUT_SESSION_ID}`
    const cancelUrl = `${process.env.NEXT_URL}`

    const line_items = Object.values(cartDetails).map((item: any) => ({
      price_data: {
        currency: 'brl',
        product_data: {
          name: item.name,
          images: [item.image],

          metadata: { productId: item.id },
        },
        unit_amount: item.price,
      },
      quantity: item.quantity,
    }))

    const checkoutSession = await stripe.checkout.sessions.create({
      success_url: successUrl,
      cancel_url: cancelUrl,
      mode: 'payment',
      line_items: line_items,
    })

    return res.status(201).json({
      checkoutUrl: checkoutSession.url,
    })
  } catch (err: any) {
    console.error('STRIPE ERROR:', err)
    return res.status(500).json({ error: err.message })
  }
}
