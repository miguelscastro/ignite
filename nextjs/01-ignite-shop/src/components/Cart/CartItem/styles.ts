import { styled } from '@/styles/index'

export const CartItemContainer = styled('div', {
  display: 'flex',
  gap: '1.25rem',
})
export const ImageContainer = styled('div', {
  background: 'linear-gradient(180deg, #1ea483 0%, #7465d4 100%)',
  borderRadius: 8,
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
  minWidth: '6.25rem',
  minHeight: '1.4375rem',

  img: {
    objectFit: 'cover',
  },
})
export const ProductInfo = styled('div', {
  display: 'flex',
  flexDirection: 'column',
  gap: '0.5rem',

  span: {
    fontWeight: 'bold',
    fontSize: '$md',
  },

  button: {
    background: 'none',
    border: 'none',
    padding: 0,
    outline: 'none',
    cursor: 'pointer',
    font: 'inherit',
    appearance: 'none',
    width: '4.0625rem',
    color: '$green300',

    '&:hover': {
      color: '$green500',
    },
  },
})

export const ProductControls = styled('div', {
  display: 'flex',

  span: {
    fontSize: '$sm',
  },
})
