import { render, screen, fireEvent } from "@testing-library/react"
import { AddUser } from "./AddUser"

describe('addUser Component', () => {
    it('update state when input values change', () => {
        render(<AddUser />)
        const studentIdInput = screen.getByLabelText('学籍番号:') as HTMLInputElement;
        const passwordInput = screen.getByLabelText('password:') as HTMLInputElement;

        fireEvent.change(studentIdInput, { target: { value: 'testUser' } });
        expect(studentIdInput.value).toBe('testUser');

        fireEvent.change(passwordInput, { target: { value: 'password' } });
        expect(passwordInput.value).toBe('password');
    })
})