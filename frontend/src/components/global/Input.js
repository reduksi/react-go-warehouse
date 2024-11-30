import React from 'react'
import { mc } from '../../lib/mc'

export default function InputText({
  value,
  name,
  label,
  onChange,
  className,
  required,
  maxLength,
  minLength,
  pattern,
  type = 'text',
  readonly,
}) {
  return (
    <main className={mc(`relative`, className)}>
      <input
        type={type}
        className={`peer p-4 block w-full border border-neutral-300 rounded-lg text-sm placeholder-opacity-0 focus:outline-none focus:border-secondary-500 focus:ring-secondary-500 disabled:opacity-50 disabled:pointer-events-none
                focus:pt-6
                focus:pb-2
                [&:not(:placeholder-shown)]:pt-6
                [&:not(:placeholder-shown)]:pb-2
                autofill:pt-6
                autofill:pb-2 ${
          readonly && 'bg-neutral-100 cursor-not-allowed'
        }`}
        placeholder=""
        value={value}
        name={name}
        onChange={onChange}
        required={required}
        maxLength={maxLength}
        minLength={minLength}
        pattern={pattern}
        readOnly={readonly}
      />
      <label
        className={`absolute top-0 start-0 p-4 text-neutral-600 h-full text-sm truncate pointer-events-none transition ease-in-out duration-100 origin-[0_0] peer-disabled:opacity-50 peer-disabled:pointer-events-none
                peer-focus:scale-90
                peer-focus:translate-x-0.5
                peer-focus:-translate-y-0.5
                peer-focus:text-xs
                peer-focus:text-neutral-500
                peer-[:not(:placeholder-shown)]:scale-90
                peer-[:not(:placeholder-shown)]:text-xs
                peer-[:not(:placeholder-shown)]:translate-x-0.5
                peer-[:not(:placeholder-shown)]:-translate-y-0.5
                peer-[:not(:placeholder-shown)]:text-neutral-500 `}
      >
        {label}
        {required && <span className="text-danger-500 ml-1 -mt-1">*</span>}
      </label>
    </main>
  )
}
