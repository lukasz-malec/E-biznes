describe('Nawigacja', () => {

    beforeEach(() => {
        cy.visit('http://localhost:5173')
    })

    it('wyświetla logo', () => {
        cy.get('.nav-logo').should('be.visible')
    })


    it('wyświetla linki nawigacji', () => {
        cy.get('.nav-links').should('be.visible')
    })



    it('przechodzi do koszyka', () => {
        cy.get('.nav-links').contains('Koszyk').click()
        cy.url().should('include', '/koszyk')
    })


    it('przechodzi do płatności', () => {
        cy.get('.nav-links').contains('Płatności').click()
        cy.url().should('include', '/platnosci')
    })

    
    it('wraca do produktów przez logo', () => {
        cy.get('.nav-links').contains('Koszyk').click()
        cy.get('.nav-logo').click()
        cy.url().should('eq', 'http://localhost:5173/')
    })

})