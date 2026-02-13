import { styled } from '@/styles/index'

export const CartContainer = styled('aside', {
  position: 'fixed',
  top: 0,
  right: 0,
  bottom: 0,
  width: '480px',
  background: '$gray800',
  display: 'flex',
  flexDirection: 'column',
  padding: '1.5rem 3rem 3rem 3rem',
  boxShadow: '-4px 0px 30px rgba(0, 0, 0, 0.8)',
  transition: 'all 0.2s ease-in-out',
  zIndex: 10,

  variants: {
    state: {
      open: {
        transform: 'translateX(0%)',
        opacity: 1,
      },
      closed: {
        transform: 'translateX(100%)',
        opacity: 0,
        pointerEvents: 'none',
      },
    },
  },
})

export const CloseCartButton = styled('button', {
  background: 'none',
  border: 'none',
  padding: 0,
  outline: 'none',
  cursor: 'pointer',
  font: 'inherit',
  color: 'inherit',
  appearance: 'none',

  width: '1rem',
  alignSelf: 'end',
  margin: '1.5rem 0 1.5rem 0',

  defaultVariants: {
    open: 'true',
  },
})
export const ItemsListContainer = styled('div', {
  display: 'flex',
  flexDirection: 'column',
  gap: '1.5rem',
  h2: {
    marginBottom: '2rem',
  },
})
export const Footer = styled('div', {
  marginTop: 'auto',
  display: 'flex',
  flexDirection: 'column',
})
export const TotalContainer = styled('div', {
  display: 'flex',
  flexDirection: 'column',
  marginBottom: '3.4375rem',
  gap: '1rem',

  div: {
    display: 'flex',
    justifyContent: 'space-between',

    '&:not(:first-child)': {
      fontWeight: 'bold',

      span: {
        fontSize: '$md',
      },
    },
  },
})
export const PurchaseButton = styled('button', {
  marginTop: 'auto',
  backgroundColor: '$green500',
  border: 0,
  color: '$white',
  borderRadius: 8,
  padding: '1.5rem',
  cursor: 'pointer',
  fontWeight: 'bold',
  fontSize: '$md',

  '&:disabled': {
    opacity: 0.6,
    cursor: 'not-allowed',
  },

  '&:not(:disabled):hover': {
    backgroundColor: '$green300',
  },
})
