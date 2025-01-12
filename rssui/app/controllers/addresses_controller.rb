class AddressesController < ApplicationController

  def index
    @addresses = Address.all
  end

  def show
    @address = Address.find(params[:id])
  end

  def new
    @address = Address.new
  end

  def destroy
    @address = Address.find(params[:id])
    if @address.delete
      redirect_to addresses_path
    end
  end

  def create
    @address = Address.new(address_params)
    if @address.save
      redirect_to addresses_path
    else
      render :new, status: :unprocessable_entity
    end
  end

  private
    def address_params
      params.expect(address: [ :url ])
    end

end
