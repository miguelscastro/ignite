import { styled } from '..'

export const Container = styled('div', {
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'flex-start',
  justifyContent: 'center',
  minHeight: '100vh',
})

export const Header = styled('header', {
  padding: '2rem 0',
  width: '100%',
  maxWidth: 1180,
  margin: '0 auto',

  display: 'flex',
  alignItems: 'center',
  justifyContent: 'space-between', // Logo na esquerda, Sacola na direita

  img: {
    cursor: 'pointer',
  },
})

export const OpenCartButton = styled('button', {
  background: '$gray800',
  border: 'none',
  borderRadius: 6,
  padding: '0.75rem',
  color: '$gray300',
  cursor: 'pointer',
  position: 'relative',

  '&:hover': {
    color: '$gray100',
  },
})
