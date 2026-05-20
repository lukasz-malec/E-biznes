// testy walidacja formularz

describe('Walidacja formularza płatności', () => {

    beforeEach(() => {
        cy.visit('http://localhost:5173')
        cy.get('.produkt-card', { timeout: 10000 }).should('have.length', 5)
        cy.get('.produkt-btn').first().click()
        cy.get('.nav-links').contains('Płatności').click()
    })


    it('formularz ma trzy pola', () => {
        cy.get('.platnosci-input').should('have.length', 3)
    })


    it('pole imię i nazwisko jest wymagane', () => {
        cy.get('[name="email"]').type('jan@test.pl')
        cy.get('[name="address"]').type('ul. Testowa 1')
        cy.get('.platnosci-btn').click()
        cy.get('[name="name"]:invalid').should('exist')
    })


    it('pole email jest wymagane', () => {
        cy.get('[name="name"]').type('Jan Kowalski')
        cy.get('[name="address"]').type('ul. Testowa 1')
        cy.get('.platnosci-btn').click()
        cy.get('[name="email"]:invalid').should('exist')
    })
    

    it('pole adres jest wymagane', () => {
        cy.get('[name="name"]').type('Jan Kowalski')
        cy.get('[name="email"]').type('jan@test.pl')
        cy.get('.platnosci-btn').click()
        cy.get('[name="address"]:invalid').should('exist')
    })

    it('pole email odrzuca nieprawidłowy format', () => {
        cy.get('[name="name"]').type('Jan Kowalski')
        cy.get('[name="email"]').type('niepoprawnyelmail')
        cy.get('[name="address"]').type('ul. Testowa 1')
        cy.get('.platnosci-btn').click()
        cy.get('[name="email"]:invalid').should('exist')
    })

    it('formularz nie wysyła się gdy wszystkie pola puste', () => {
        cy.get('.platnosci-btn').click()
        cy.get('.platnosci-input:invalid').should('exist')
    })


    it('pola przyjmują wpisany tekst', () => {
        cy.get('[name="name"]').type('Jan Kowalski')
        cy.get('[name="name"]').should('have.value', 'Jan Kowalski')
    })



    it('pole email przyjmuje poprawny format', () => {
        cy.get('[name="email"]').type('jan@test.pl')
        cy.get('[name="email"]').should('have.value', 'jan@test.pl')
    })


    it('przycisk złóż zamówienie jest widoczny', () => {
        cy.get('.platnosci-btn').should('be.visible')
    })


    it('poprawnie wypełniony formularz wysyła zamówienie', () => {
        cy.get('[name="name"]').type('Jan Kowalski')
        cy.get('[name="email"]').type('jan@test.pl')
        cy.get('[name="address"]').type('ul. Testowa 1, Kraków')
        cy.get('.platnosci-btn').click()
        cy.contains('Zamówienie złożone', { timeout: 10000 }).should('be.visible')
    })

})