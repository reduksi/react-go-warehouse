import { useRef, useCallback } from 'react'
import { AiOutlineCalendar } from 'react-icons/ai';
import { mc } from '../../lib/mc'

const DateInput = ({
  value,
  name,
  label,
  onChange,
  className,
  required,
  readonly,
}) => {
  const inputRef = useRef(null)

  const handleIconClick = (e) => {
    e.stopPropagation()
    const dateInput = inputRef.current
    if (dateInput) {
      dateInput.focus()
      dateInput.showPicker()
    }
  }

  const handleInputMouseDown = useCallback(
    (e) => {
      e.stopPropagation()
      const dateInput = inputRef.current
      if (dateInput) {
        dateInput.focus()
        dateInput.showPicker()
      }
    },
    []
  )

  return (
    <div className={mc('relative', className)}>
      <input
        type="date"
        id="date"
        name={name}
        required={required}
        className={`peer w-full border border-neutral-300 rounded-lg pt-6 pb-2 px-2 text-sm text-gray-900 placeholder-transparent focus:outline-none focus:ring-1 focus:ring-primary-500 custom-date-input ${
          readonly && 'bg-neutral-100 cursor-not-allowed'
        }`}
        value={value}
        onChange={onChange}
        ref={inputRef}
        onMouseDown={handleInputMouseDown} // Handle mouse down on input
      />
      <label
        htmlFor="date"
        className={`absolute left-2.5 top-2 transition-all duration-200 ease-in-out text-neutral-500 text-xs translate-y-0.3 pointer-events-none`}
      >
        {label}
        {required && <span className="text-danger-500 ml-1 -mt-1">*</span>}
      </label>
      <AiOutlineCalendar
        className="absolute right-3 top-1/2 transform -translate-y-1/2 h-5 w-5 text-neutral-500 cursor-pointer"
        onClick={handleIconClick} // Handle calendar icon click
      />
      <style jsx>{`
        .custom-date-input::-webkit-calendar-picker-indicator {
          opacity: 0;
        }
      `}</style>
    </div>
  )
}

export default DateInput
