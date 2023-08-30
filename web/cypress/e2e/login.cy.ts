describe('template spec', () => {
  it('Visits the Hexa|Gym', () => {
    cy.visit('http://localhost:3000/login')
    cy.url().should('include', '/login')
    cy.get('h2').should('contain', 'Welcome to Hexa|Gym expert!')
  })

  it('login to Hexa|Gym', () => {
    cy.visit('http://localhost:3000/login')
    cy.get('input[name=username]').type('assembly')
    cy.get('input[name=password]').type('1234')
    cy.get('button[type=submit]').click()
    cy.get('strong').should('contain', 'Hexa|Gym expert')
    cy.get('strong').should('contain', 'Hexa|Gym expert')
  })
})
