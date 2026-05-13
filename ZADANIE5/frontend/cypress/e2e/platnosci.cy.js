
describe('Płatności', () => {

    beforeEach(() => {
        cy.visit('http://localhost:5173')
        cy.get('.produkt-card', { timeout: 10000 }).should('have.length', 5)
    })

    it('wyświetla formularz płatności', () => {
        cy.get('.nav-links').contains('Płatności').click()
        cy.get('.platnosci-input').should('have.length', 3)
    })


    it('nie wysyła formularza bez danych', () => {
        cy.get('.nav-links').contains('Płatności').click()
        cy.get('.platnosci-btn').click()
        cy.get('.platnosci-input:invalid').should('exist')
    })

    
    it('wypełnia i wysyła formularz', () => {
        cy.get('.produkt-btn').first().click()
        cy.get('.nav-links').contains('Płatności').click()
        cy.get('[name="name"]').type('Jan Kowalski')
        cy.get('[name="email"]').type('jan@test.pl')
        cy.get('[name="address"]').type('ul. Testowa 1, Kraków')
        cy.get('.platnosci-btn').click()
        cy.contains('Zamówienie złożone', { timeout: 10000 }).should('be.visible')
    })



})